#include <QApplication>
#include <QPushButton>
#include <iostream>
#include <QFontDatabase>

#include "grpcpp/grpcpp.h"
#include "Service/Service.grpc.pb.h"
#include "src/styles/QtAdvancedStylesheet.h"
#include "src/ui/mainwindow.h"
#include "src/ui/LoginWindow.h"
#include "src/framework/service/RpcCall.h"
#include "src/framework/config/AppConfig.h"

int main(int argc, char *argv[]) {
    setbuf(stdout, nullptr);



    QApplication a(argc, argv);

    QString fontPath = QString(R"(C:\Users\KiraY\Desktop\GoTest\src\SwcManagerClient\Resource\fonts\SourceHanSansCN\SourceHanSansCN-Regular.ttf)");
    int loadedFontID = QFontDatabase::addApplicationFont(fontPath);
    QStringList loadedFontFamilies = QFontDatabase::applicationFontFamilies(loadedFontID);
    if (!loadedFontFamilies.empty()) {
        const QString& sansCNFamily = loadedFontFamilies.at(0);
        QFont defaultFont = QApplication::font();
        defaultFont.setFamily(sansCNFamily);
        defaultFont.setPixelSize(14);
        QApplication::setFont(defaultFont);
    }

    QString appDir = qApp->applicationDirPath();
    acss::QtAdvancedStylesheet styleManager;
    styleManager.setStylesDirPath(R"(C:\Users\KiraY\Desktop\GoTest\src\SwcManagerClient\Resource\styles)");
    styleManager.setOutputDirPath(appDir + "/StylesOutput");
    styleManager.setCurrentStyle("qt_material_modified");
    styleManager.setCurrentTheme("light_blue");
    styleManager.updateStylesheet();
    qApp->setStyleSheet(styleManager.styleSheet());
    // setWindowIcon(advancedStyleSheet.styleIcon());
    // qApp->setStyleSheet(advancedStyleSheet.styleSheet());
    // connect(&advancedStyleSheet, SIGNAL(stylesheetChanged()), this,
    //     SLOT(onStyleManagerStylesheetChanged()));

    AppConfig::getInstance().initialize("AppConfig.json");
    AppConfig::getInstance().readConfig();

    RpcCall::getInstance().initialize("127.0.0.1:8080");

    LoginWindow loginWindow{};
    if(loginWindow.exec() != QDialog::Accepted) {
        return -1;
    }

    MainWindow mainwindow{};
    mainwindow.resize(1920/2, 1080/2);
    mainwindow.show();

    return QApplication::exec();
}
