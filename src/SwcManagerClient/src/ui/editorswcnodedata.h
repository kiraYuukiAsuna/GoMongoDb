#pragma once

#include <QWidget>
#include "../Generated/Message/Message.pb.h"
#include "EditorBase.h"


QT_BEGIN_NAMESPACE
namespace Ui { class EditorSwcNodeData; }
QT_END_NAMESPACE

class EditorSwcNodeData : public QWidget, public EditorBase{
Q_OBJECT

public:
    explicit EditorSwcNodeData(const std::string& swcName, QWidget *parent = nullptr);
    ~EditorSwcNodeData() override;

    virtual std::string getName() {
        return "";
    }

    virtual MetaInfoType getMetaInfoType() {
        return MetaInfoType::eSwcData;
    }
private:
    void getSwcNodeData();

    proto::SwcMetaInfoV1 m_SwcMetaInfo;
    std::string m_SwcName;

    Ui::EditorSwcNodeData *ui;
};

