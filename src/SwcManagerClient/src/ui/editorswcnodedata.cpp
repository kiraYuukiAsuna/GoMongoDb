//
// Created by KiraY on 2023/11/22.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorSwcNodeData.h" resolved

#include "editorswcnodedata.h"
#include "ui_EditorSwcNodeData.h"


EditorSwcNodeData::EditorSwcNodeData(QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorSwcNodeData) {
    ui->setupUi(this);
}

EditorSwcNodeData::~EditorSwcNodeData() {
    delete ui;
}
