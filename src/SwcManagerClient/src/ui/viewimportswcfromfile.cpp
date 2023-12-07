#include "viewimportswcfromfile.h"
#include "ui_ViewImportSwcFromFile.h"
#include <QFileDialog>
#include <QStandardPaths>
#include "src/swcio/SwcReader.h"

ViewImportSwcFromFile::ViewImportSwcFromFile(QWidget *parent) :
    QDialog(parent), ui(new Ui::ViewImportSwcFromFile) {
    ui->setupUi(this);

    ui->SwcFileInfo->clear();
    ui->SwcFileInfo->setColumnCount(3);
    QStringList headerLabels;
    headerLabels
            << "Swc FilePath"
            << "Type"
            << "Detected Swc Node Number";
    ui->SwcFileInfo->setHorizontalHeaderLabels(headerLabels);

    connect(ui->SelectBtn,&QPushButton::clicked,this,[this](){
        QFileDialog fileDialog(this);
        fileDialog.setWindowTitle("选择Swc文件");
        fileDialog.setDirectory(QStandardPaths::writableLocation(QStandardPaths::HomeLocation));
        fileDialog.setNameFilter(tr("File(*.swc *.eswc *.*)"));
        fileDialog.setFileMode(QFileDialog::ExistingFiles);
        fileDialog.setViewMode(QFileDialog::Detail);

        QStringList fileNames;
        if (fileDialog.exec()) {
            fileNames = fileDialog.selectedFiles();
            ui->SwcFileInfo->setRowCount(fileNames.size());
            for(int i=0; i< fileNames.size(); i++){
                std::filesystem::path filePath(fileNames[0].toStdString());
                if(filePath.extension() == ".swc"){
                    Swc swc(filePath.string());
                    auto neuron = swc.getNeuron();

                    ui->SwcFileInfo->setItem(i, 0,
                                             new QTableWidgetItem(QString::fromStdString(std::to_string(1))));
                    ui->SwcFileInfo->setItem(i, 1,
                                             new QTableWidgetItem(QString::fromStdString(std::to_string(1))));
                    ui->SwcFileInfo->setItem(i, 2,
                                             new QTableWidgetItem(QString::fromStdString(std::to_string(1))));
                    
                }else if(filePath.extension() == ".eswc"){
                    ESwc eSwc(filePath.string());
                    auto neuron = eSwc.getNeuron();


                }
            }
        }

    });

    connect(ui->ImportBtn,&QPushButton::clicked,this,[this](){

    });

    connect(ui->CancelBtn,&QPushButton::clicked,this,[this](){

    });
}

ViewImportSwcFromFile::~ViewImportSwcFromFile() {
    delete ui;
}
