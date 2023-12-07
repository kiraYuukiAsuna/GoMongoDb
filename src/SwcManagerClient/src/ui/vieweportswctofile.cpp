#include "vieweportswctofile.h"
#include "ui_ViewEportSwcToFile.h"


ViewEportSwcToFile::ViewEportSwcToFile(QWidget *parent) :
    QDialog(parent), ui(new Ui::ViewEportSwcToFile) {
    ui->setupUi(this);
}

ViewEportSwcToFile::~ViewEportSwcToFile() {
    delete ui;
}
