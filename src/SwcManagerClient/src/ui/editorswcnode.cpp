//
// Created by KiraY on 2023/11/23.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorSwcNode.h" resolved

#include "editorswcnode.h"
#include "ui_EditorSwcNode.h"


EditorSwcNode::EditorSwcNode(QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorSwcNode) {
    ui->setupUi(this);
}

EditorSwcNode::~EditorSwcNode() {
    delete ui;
}
