//
// Created by KiraY on 2023/11/20.
//

// You may need to build the project (run Qt uic code generator) to get "ui_RegisterWindow.h" resolved

#include "registerwindow.h"

#include <QMessageBox>
#include <grpcpp/client_context.h>
#include <Message/Request.pb.h>

#include "ui_RegisterWindow.h"
#include "src/framework/service/RpcCall.h"


RegisterWindow::RegisterWindow(QWidget *parent) :
    QDialog(parent), ui(new Ui::RegisterWindow) {
    ui->setupUi(this);

    this->setWindowFlags(this->windowFlags() & ~Qt::WindowMaximizeButtonHint);
    this->setFixedSize(this->size());

    connect(ui->registerBtn,&QPushButton::clicked, this, &RegisterWindow::onRegisterBtnClicked);


}

RegisterWindow::~RegisterWindow() {
    delete ui;
}

void RegisterWindow::onRegisterBtnClicked(bool checked) {
    if(ui->userNameEditor->text().isEmpty()) {
        QMessageBox::warning(this,"Error","User Name cannot be empty!");
        return;
    }
    if(ui->passwordEditor->text().isEmpty() || ui->repeatPasswordEditor->text().isEmpty()) {
        QMessageBox::warning(this,"Error","Password cannot be empty!");
        return;
    }
    if(ui->passwordEditor->text().trimmed() != ui->repeatPasswordEditor->text().trimmed()) {
        QMessageBox::warning(this,"Error","Password and Repeated Password are not equal!");
        return;
    }

    grpc::ClientContext context;
    proto::CreateUserRequest request;
    request.mutable_userinfo()->set_name(ui->userNameEditor->text().trimmed().toStdString());
    request.mutable_userinfo()->set_password(ui->repeatPasswordEditor->text().trimmed().toStdString());

    auto& rpcCall = RpcCall::getInstance();
    proto::CreateUserResponse response;
    auto status = rpcCall.Stub()->CreateUser(&context, request, &response);
    if(status.ok()){
        if(response.status()) {
            QMessageBox::information(this,"Info","Register Successfully!");
            accept();
        }else {
            QMessageBox::warning(this,"Info","Register Failed!" + QString::fromStdString(response.message()));
        }

    }else{
        QMessageBox::critical(this,"Error",QString::fromStdString(status.error_message()));
    }
}
