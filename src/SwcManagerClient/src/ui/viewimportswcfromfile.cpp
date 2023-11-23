//
// Created by KiraY on 2023/11/23.
//

// You may need to build the project (run Qt uic code generator) to get "ui_ViewImportSwcFromFile.h" resolved

#include "viewimportswcfromfile.h"
#include "ui_ViewImportSwcFromFile.h"


ViewImportSwcFromFile::ViewImportSwcFromFile(QWidget *parent) :
    QDialog(parent), ui(new Ui::ViewImportSwcFromFile) {
    ui->setupUi(this);
}

ViewImportSwcFromFile::~ViewImportSwcFromFile() {
    delete ui;
}
