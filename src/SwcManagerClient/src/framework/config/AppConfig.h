#pragma once

#include <fstream>
#include <QStandardPaths>
#include <string>
#include "json.hpp"

class AppConfig {
public:
    enum class ConfigItem {
        eCachedUserName = 0,
        eCachedPassword,
        eAccountExpiredTime
    };

    static AppConfig& getInstance() {
        static AppConfig instance;
        return instance;
    }

    void initialize(const std::string& appConfigFile) {
        m_AppConfigFile = appConfigFile;

        std::filesystem::path path(m_AppRoamingPath);
        if(!std::filesystem::exists(path)) {
            std::filesystem::create_directories(path);
        }
    }

    void readConfig() {
        std::filesystem::path path(m_AppRoamingPath);
        path = path / m_AppConfigFile;
        std::ifstream f(path.string());
        try {
            m_AppConfig = nlohmann::json::parse(f);
        }catch (...) {

        }
        f.close();
    }

    void writeConfig() {
        std::filesystem::path path(m_AppRoamingPath);
        path = path / m_AppConfigFile;
        std::ofstream f(path.string());
        f<<m_AppConfig;
        f.close();
    }

    std::string getConfig(ConfigItem configItem) {
        switch (configItem) {
            case ConfigItem::eCachedUserName: {
                if(m_AppConfig.contains("eCachedUserName")) {
                    return m_AppConfig["eCachedUserName"];
                }else {
                    return "";
                }
                break;
            }
            case ConfigItem::eCachedPassword: {
                if(m_AppConfig.contains("eCachedPassword")) {
                    return m_AppConfig["eCachedPassword"];
                }else {
                    return "";
                }
                break;
            }
            case ConfigItem::eAccountExpiredTime: {
                if(m_AppConfig.contains("eAccountExpiredTime")) {
                    return m_AppConfig["eAccountExpiredTime"];
                }else {
                    return "";
                }
                break;
            }
            default: {
                return "ConfigError";
                break;
            }
        }
    }

    void setConfig(ConfigItem configItem, const std::string& configData) {
        switch (configItem) {
            case ConfigItem::eCachedUserName: {
                m_AppConfig["eCachedUserName"] = configData;
                break;
            }
            case ConfigItem::eCachedPassword: {
                m_AppConfig["eCachedPassword"] = configData;
                break;
            }
            case ConfigItem::eAccountExpiredTime: {
                m_AppConfig["eAccountExpiredTime"] = configData;
                break;
            }
            default: {
                break;
            }
        }
    }

private:
    AppConfig() {
        m_AppRoamingPath = QStandardPaths::writableLocation(QStandardPaths::AppDataLocation).toStdString();
    }

    std::string m_AppRoamingPath;

    std::string m_AppConfigFile;
    nlohmann::json m_AppConfig;
};

