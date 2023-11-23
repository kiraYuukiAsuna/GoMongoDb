//
// Created by KiraY on 2023/11/23.
//

// You may need to build the project (run Qt uic code generator) to get "ui_ViewEportSwcToFile.h" resolved

#include "vieweportswctofile.h"
#include "ui_ViewEportSwcToFile.h"


ViewEportSwcToFile::ViewEportSwcToFile(QWidget *parent) :
    QDialog(parent), ui(new Ui::ViewEportSwcToFile) {
    ui->setupUi(this);
}

ViewEportSwcToFile::~ViewEportSwcToFile() {
    delete ui;
}
