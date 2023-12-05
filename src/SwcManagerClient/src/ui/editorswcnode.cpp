#include "editorswcnode.h"
#include "ui_EditorSwcNode.h"


EditorSwcNode::EditorSwcNode(proto::GetSwcNodeDataResponse& response, QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorSwcNode) {
    ui->setupUi(this);

    refresh(response);
}

EditorSwcNode::~EditorSwcNode() {
    delete ui;
}

void EditorSwcNode::refresh(proto::GetSwcNodeDataResponse& response) {
    for(int i=0;i<response.swcnodedata().swcdata_size();i++) {

    }
}
