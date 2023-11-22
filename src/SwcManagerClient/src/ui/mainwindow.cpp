#include "MainWindow.h"
#include "ui_mainwindow.h"


MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent), ui(new Ui::MainWindow) {
    ui->setupUi(this);
    setWindowState(Qt::WindowMaximized);

    m_Splitter = new QSplitter(this);

    m_LeftClientView = new LeftClientView(this);
    m_RightClientView = new RightClientView(this);

    m_Splitter->addWidget(m_LeftClientView);
    m_Splitter->addWidget(m_RightClientView);
    m_Splitter->setSizes(QList<int>()<<100000000<<400000000);
    m_Splitter->setCollapsible(0, false);
    m_Splitter->setCollapsible(1, false);

    this->setCentralWidget(m_Splitter);

    auto menuBar = new QMenuBar(this);

    auto* menuFile = new QMenu(menuBar);
    // icon
    menuFile->setTitle("File");
    menuBar->addMenu(menuFile);

    auto* menuImportSwcFile = new QAction(menuFile);
    menuImportSwcFile->setText("Import Swc File");
    menuFile->addAction(menuImportSwcFile);

    auto* menuExportToSwcFile = new QAction(menuFile);
    menuExportToSwcFile->setText("Import Swc File");
    menuFile->addAction(menuExportToSwcFile);

    setMenuBar(menuBar);
}

MainWindow::~MainWindow() {
    delete ui;
}

LeftClientView &MainWindow::getLeftClientView() {
    return *m_LeftClientView;
}

RightClientView &MainWindow::getRightClientView() {
    return *m_RightClientView;
}
