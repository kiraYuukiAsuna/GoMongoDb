//
// Created by KiraY on 2023/11/18.
//

// You may need to build the project (run Qt uic code generator) to get "ui_mainwindow.h" resolved

#include "MainWindow.h"
#include "ui_mainwindow.h"


MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent), ui(new Ui::MainWindow) {
    ui->setupUi(this);

}

MainWindow::~MainWindow() {
    delete ui;
}
