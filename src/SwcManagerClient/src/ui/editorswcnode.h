#pragma once

#include <QWidget>
#include <Message/Request.pb.h>
#include <Message/Response.pb.h>


QT_BEGIN_NAMESPACE
namespace Ui { class EditorSwcNode; }
QT_END_NAMESPACE

class EditorSwcNode : public QWidget {
Q_OBJECT

public:
    explicit EditorSwcNode(proto::GetSwcNodeDataResponse& response, QWidget *parent = nullptr);
    ~EditorSwcNode() override;

    void refresh(proto::GetSwcNodeDataResponse& response);

private:
    Ui::EditorSwcNode *ui;
};
