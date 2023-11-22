#include "leftclientview.h"
#include "ui_LeftClientView.h"


LeftClientView::LeftClientView(QWidget *parent) :
    QWidget(parent), ui(new Ui::LeftClientView) {
    ui->setupUi(this);
}

LeftClientView::~LeftClientView() {
    delete ui;
}
