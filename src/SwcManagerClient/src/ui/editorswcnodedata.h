//
// Created by KiraY on 2023/11/22.
//

#ifndef EDITORSWCNODEDATA_H
#define EDITORSWCNODEDATA_H

#include <QWidget>

#include "EditorBase.h"


QT_BEGIN_NAMESPACE
namespace Ui { class EditorSwcNodeData; }
QT_END_NAMESPACE

class EditorSwcNodeData : public QWidget, public EditorBase{
Q_OBJECT

public:
    explicit EditorSwcNodeData(QWidget *parent = nullptr);
    ~EditorSwcNodeData() override;

    virtual std::string getName() {
        return "";
    }

    virtual MetaInfoType getMetaInfoType() {
        return MetaInfoType::eSwcData;
    }
private:
    Ui::EditorSwcNodeData *ui;
};


#endif //EDITORSWCNODEDATA_H
