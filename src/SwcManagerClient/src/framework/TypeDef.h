#pragma once
#include <QObject>

enum class MetaInfoType{
    eProjectContainer = 0,
    eProject,
    eSwcContainer,
    eSwc,
    eDailyStatisticsContainer,
    eDailyStatistics
};

struct LeftClientViewTreeWidgetMetaInfo{
    MetaInfoType type;
    std::string name;
};
Q_DECLARE_METATYPE(LeftClientViewTreeWidgetMetaInfo)
