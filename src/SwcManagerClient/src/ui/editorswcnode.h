//
// Created by KiraY on 2023/11/23.
//

#ifndef EDITORSWCNODE_H
#define EDITORSWCNODE_H

#include <QWidget>


QT_BEGIN_NAMESPACE
namespace Ui { class EditorSwcNode; }
QT_END_NAMESPACE

class EditorSwcNode : public QWidget {
Q_OBJECT

public:
    explicit EditorSwcNode(QWidget *parent = nullptr);
    ~EditorSwcNode() override;

private:
    Ui::EditorSwcNode *ui;
};


#endif //EDITORSWCNODE_H
