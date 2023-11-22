#include "rightclientview.h"
#include "ui_RightClientView.h"


RightClientView::RightClientView(QWidget *parent) :
    QWidget(parent), ui(new Ui::RightClientView) {
    ui->setupUi(this);
}

RightClientView::~RightClientView() {
    delete ui;
}
