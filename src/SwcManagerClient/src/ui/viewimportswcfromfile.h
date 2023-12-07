#pragma once

#include <QDialog>
#include "src/swcio/SwcReader.h"


QT_BEGIN_NAMESPACE
namespace Ui { class ViewImportSwcFromFile; }
QT_END_NAMESPACE

class ViewImportSwcFromFile : public QDialog {
Q_OBJECT

public:
    explicit ViewImportSwcFromFile(QWidget *parent = nullptr);
    ~ViewImportSwcFromFile() override;

private:
    Ui::ViewImportSwcFromFile *ui;

    std::vector<Swc> m_SwcList;
    std::vector<ESwc> m_ESwcList;
    bool m_ActionImportComplete= false;
};
