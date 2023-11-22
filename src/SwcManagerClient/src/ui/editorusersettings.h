//
// Created by KiraY on 2023/11/22.
//

#ifndef EDITORUSERSETTINGS_H
#define EDITORUSERSETTINGS_H

#include <QWidget>

#include "EditorBase.h"


QT_BEGIN_NAMESPACE
namespace Ui { class EditorUserSettings; }
QT_END_NAMESPACE

class EditorUserSettings : public QWidget,public EditorBase{
Q_OBJECT

public:
    explicit EditorUserSettings(QWidget *parent = nullptr);
    ~EditorUserSettings() override;

    virtual std::string getName() {
        return "";
    }

    virtual MetaInfoType getMetaInfoType() {
        return MetaInfoType::eUserMetaInfo;
    }
private:
    Ui::EditorUserSettings *ui;
};


#endif //EDITORUSERSETTINGS_H
