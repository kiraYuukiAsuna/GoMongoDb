#pragma once

#include <QWidget>
#include <QDialog>
#include "EditorBase.h"

QT_BEGIN_NAMESPACE
namespace Ui { class EditorUserSettings; }
QT_END_NAMESPACE

class LeftClientView;

class EditorUserSettings : public QDialog, public EditorBase{
Q_OBJECT

public:
    explicit EditorUserSettings(LeftClientView *leftClientView);
    ~EditorUserSettings() override;

    virtual std::string getName() {
        return "EditorUserSettings";
    }

    virtual MetaInfoType getMetaInfoType() {
        return MetaInfoType::eUserMetaInfo;
    }

    void getUserMetaInfo();

private:
    Ui::EditorUserSettings *ui;
    LeftClientView* m_LeftClientView;
};
