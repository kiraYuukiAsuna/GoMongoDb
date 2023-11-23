//
// Created by KiraY on 2023/11/23.
//

// You may need to build the project (run Qt uic code generator) to get "ui_ViewCreateSwc.h" resolved

#include "viewcreateswc.h"
#include "ui_ViewCreateSwc.h"


ViewCreateSwc::ViewCreateSwc(QWidget *parent) :
    QWidget(parent), ui(new Ui::ViewCreateSwc) {
    ui->setupUi(this);
}

ViewCreateSwc::~ViewCreateSwc() {
    delete ui;
}
