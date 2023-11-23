//
// Created by KiraY on 2023/11/23.
//

#ifndef VIEWIMPORTSWCFROMFILE_H
#define VIEWIMPORTSWCFROMFILE_H

#include <QDialog>


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
};


#endif //VIEWIMPORTSWCFROMFILE_H
