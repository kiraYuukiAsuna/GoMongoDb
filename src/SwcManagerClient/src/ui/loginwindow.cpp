#include "loginwindow.h"
#include "ui_LoginWindow.h"
#include <QMessageBox>
#include <Message/Request.pb.h>
#include "registerwindow.h"
#include "src/framework/service/RpcCall.h"
#include "src/framework/config/AppConfig.h"
#include <QTimer>
#include <QTimerEvent>

LoginWindow::LoginWindow(QWidget *parent) :
    QDialog(parent), ui(new Ui::LoginWindow) {
    ui->setupUi(this);

    this->setWindowFlags(this->windowFlags() & ~Qt::WindowMaximizeButtonHint);
    this->setFixedSize(this->size());

    connect(ui->loginBtn,&QPushButton::clicked, this, &LoginWindow::onLoginBtnClicked);
    connect(ui->registerBtn, &QPushButton::clicked, this, &LoginWindow::onRegisterBtnClicked);

    static QTimer timer;
    connect(&timer,&QTimer::timeout,this,[&]() {
        verifyCachedAccount();
    });
    timer.setSingleShot(true);
    timer.start(300);
}

LoginWindow::~LoginWindow() {
    delete ui;
}

void LoginWindow::verifyCachedAccount() {
    auto cachedUserName = AppConfig::getInstance().getConfig(AppConfig::ConfigItem::eCachedUserName);
    auto cachedPassword = AppConfig::getInstance().getConfig(AppConfig::ConfigItem::eCachedPassword);
    auto accountExpiredTime = AppConfig::getInstance().getConfig(AppConfig::ConfigItem::eAccountExpiredTime);

    auto timestampeNow = std::chrono::system_clock::now().time_since_epoch().count();
    long long timestampeAccountExpired = 0;
    if(!accountExpiredTime.empty()) {
        timestampeAccountExpired = std::stoll(accountExpiredTime);
    }

    if(!cachedUserName.empty() && !cachedPassword.empty() && timestampeAccountExpired > timestampeNow) {
        doLogin(QString::fromStdString(cachedUserName), QString::fromStdString(cachedPassword), true);
    }
}

void LoginWindow::onLoginBtnClicked(bool checked) {
    if(ui->userNameEditor->text().isEmpty()) {
        QMessageBox::warning(this,"Error","User Name cannot be empty!");
        return;
    }
    if(ui->passwordEditor->text().isEmpty()) {
        QMessageBox::warning(this,"Error","Password cannot be empty!");
        return;
    }
    doLogin(ui->userNameEditor->text().trimmed(), ui->passwordEditor->text().trimmed());
}

void LoginWindow::onRegisterBtnClicked(bool checked) {
    RegisterWindow registerWindow{this};
    registerWindow.exec();
}

bool LoginWindow::doLogin(QString userName, QString password, bool slientMode) {
    grpc::ClientContext context;
    proto::UserLoginRequest request;
    request.set_username(userName.toStdString());
    request.set_password(password.toStdString());

    auto& rpcCall = RpcCall::getInstance();
    proto::UserLoginResponse response;
    auto status = rpcCall.Stub()->UserLogin(&context, request, &response);
    if(status.ok()){
        if(response.status()) {
            if(!slientMode) {
                QMessageBox::information(this,"Info","Login Successfully!");
            }
            AppConfig::getInstance().setConfig(AppConfig::ConfigItem::eCachedUserName, userName.toStdString());
            AppConfig::getInstance().setConfig(AppConfig::ConfigItem::eCachedPassword, password.toStdString());

            auto timestampeNow = std::chrono::system_clock::now();
            std::chrono::seconds days(15);
            auto expiredTime = timestampeNow + days;
            auto seconds_since_epoch = expiredTime.time_since_epoch().count();

            AppConfig::getInstance().setConfig(AppConfig::ConfigItem::eAccountExpiredTime, std::to_string(seconds_since_epoch));

            AppConfig::getInstance().writeConfig();
            accept();
            return true;
        }else {
            if(!slientMode){
                QMessageBox::warning(this,"Info","Login Failed!" + QString::fromStdString(response.message()));
            }
        }

    }else{
        if(!slientMode) {
            QMessageBox::critical(this,"Error",QString::fromStdString(status.error_message()));
        }
    }
    return false;
}
