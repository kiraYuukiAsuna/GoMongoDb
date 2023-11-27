#include "editorswcnodedata.h"
#include "ui_EditorSwcNodeData.h"
#include "Message/Response.pb.h"
#include "src/framework/service/WrappedCall.h"

EditorSwcNodeData::EditorSwcNodeData(const std::string& swcName, QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorSwcNodeData) {
    ui->setupUi(this);
    m_SwcName = swcName;


}

EditorSwcNodeData::~EditorSwcNodeData() {
    delete ui;
}

void EditorSwcNodeData::getSwcNodeData() {
    proto::GetSwcMetaInfoResponse response;
    if(WrappedCall::getSwcMetaInfoByName(m_SwcName, response, this)){
        m_SwcMetaInfo.CopyFrom(response.swcinfo());
    }
}
