#include "viewcreateswc.h"

#include <QMessageBox>
#include <Message/Request.pb.h>
#include <Message/Response.pb.h>

#include "ui_ViewCreateSwc.h"
#include "src/framework/service/CachedProtoData.h"
#include "src/framework/service/RpcCall.h"

ViewCreateSwc::ViewCreateSwc(QWidget* parent) : QDialog(parent), ui(new Ui::ViewCreateSwc) {
    ui->setupUi(this);

    connect(ui->CancelBtn, &QPushButton::clicked, this, [&]() {
        reject();
    });

    connect(ui->OKBtn, &QPushButton::clicked, this, [&]() {
        proto::CreateSwcRequest request;
        proto::CreateSwcResponse response;
        grpc::ClientContext context;

        request.mutable_userinfo()->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

        if (ui->Name->text().trimmed().isEmpty()) {
            QMessageBox::warning(this, "Error", "Name cannot be empty!");
            return false;
        }
        if (ui->Description->text().isEmpty()) {
            QMessageBox::warning(this, "Error", "Description cannot be empty!");
            return false;
        }
        request.mutable_swcinfo()->set_name(ui->Name->text().toStdString());
        request.mutable_swcinfo()->set_description(ui->Description->text().toStdString());

        auto status = RpcCall::getInstance().Stub()->CreateSwc(&context, request, &response);
        if (status.ok()) {
            if (response.status()) {
                QMessageBox::information(this, "Info", "Create Swc Successfully!");
                accept();
                return true;
            }
            QMessageBox::critical(this, "Error", QString::fromStdString(response.message()));
        }
        QMessageBox::critical(this, "Error", QString::fromStdString(status.error_message()));
        return false;
    });
}

ViewCreateSwc::~ViewCreateSwc() {
    delete ui;
}
