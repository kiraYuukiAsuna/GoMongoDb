//
// Created by KiraY on 2023/11/22.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorUserSettings.h" resolved

#include "editorusersettings.h"
#include "ui_EditorUserSettings.h"


EditorUserSettings::EditorUserSettings(QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorUserSettings) {
    ui->setupUi(this);
}

EditorUserSettings::~EditorUserSettings() {
    delete ui;
}
