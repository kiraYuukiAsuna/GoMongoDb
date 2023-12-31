#pragma once

#include <string>
#include <vector>
#include <fstream>
#include <ranges>
#include <sstream>

inline std::vector<std::string> string_split(const std::string &str, char delim) {
    std::stringstream ss(str);
    std::string item;
    std::vector<std::string> elems;
    while (std::getline(ss, item, delim)) {
        if (!item.empty()) {
            elems.push_back(item);
        }
    }
    return elems;
}

struct NeuronUnit {
    NeuronUnit() {

    }

    NeuronUnit(float x, float y, float z) : x(x), y(y), z(z) {

    }

    std::string getString(bool isApoOutput = false) {
        std::string str =
                std::to_string(n) + " " +
                std::to_string(type) + " " +
                std::to_string(x) + " " +
                std::to_string(y) + " " +
                std::to_string(z) + " " +
                std::to_string(radius) + " " +
                std::to_string(parent) + " " +
                std::to_string(seg_id) + " " +
                std::to_string(level) + " " +
                std::to_string(mode) + " " +
                std::to_string(timestamp) + " " +
                std::to_string(feature_value);
        return str;
    }

    int n = 0;
    int type = 0;
    float x = 0.0;
    float y = 0.0;
    float z = 0.0;
    float radius = 0.0;
    int parent = 0;
    int seg_id = 0;
    int level = 0;
    int mode = 0;
    int timestamp = 0;
    int feature_value = 0;
};

class ESwc {
public:
    ESwc(std::string filePath)
            : m_FilePath(filePath) {
    }

    void ReadFromFile() {
        std::ifstream infile;
        infile.open(m_FilePath);
        if (!infile.is_open()) {
            return;
        }

        std::string rowContent;
        while (std::getline(infile, rowContent)) {
            auto splitResult = string_split(rowContent, ' ');

            if (rowContent.empty() || rowContent[0] == '#') {
                rowContent.clear();
                continue;
            }

            if (splitResult.size() < 12) {
                continue;
            }
            NeuronUnit unit;
            unit.n = std::stoi(splitResult[0]);
            unit.type = std::stoi(splitResult[1]);
            unit.x = std::stof(splitResult[2]);
            unit.y = std::stof(splitResult[3]);
            unit.z = std::stof(splitResult[4]);
            unit.radius = std::stof(splitResult[5]);
            unit.parent = std::stoi(splitResult[6]);
            unit.seg_id = std::stoi(splitResult[7]);
            unit.level = std::stoi(splitResult[8]);
            unit.mode = std::stoi(splitResult[9]);
            unit.timestamp = std::stoi(splitResult[10]);
            unit.feature_value = std::stoi(splitResult[11]);

            m_Neuron.push_back(unit);

            rowContent.clear();
        }
    }

    bool WriteToFile() {
        std::ofstream outfile;
        outfile.open(m_FilePath);
        if (!outfile.is_open()) {
            return false;
        }

        outfile << "# Generated by SwcManagerClient\n"
                << "# Source File(s):\n"
                << "# id,type,x,y,z,r,pid\n";

        for (auto &neuron: m_Neuron) {
            outfile << std::to_string(neuron.n) + " " +
                       std::to_string(neuron.type) + " " +
                       std::to_string(neuron.x) + " " +
                       std::to_string(neuron.y) + " " +
                       std::to_string(neuron.z) + " " +
                       std::to_string(neuron.radius) + " " +
                       std::to_string(neuron.parent) + " " +
                       std::to_string(neuron.seg_id) + " " +
                       std::to_string(neuron.level) + " " +
                       std::to_string(neuron.mode) + " " +
                       std::to_string(neuron.timestamp) + " " +
                       std::to_string(neuron.feature_value)<<"\n";
        }
        outfile.close();
        return true;
    }

    std::vector<NeuronUnit> &getNeuron() {
        return m_Neuron;
    }

    void setNeuron(std::vector<NeuronUnit>& neuron) {
        m_Neuron = neuron;
    }

private:
    std::vector<NeuronUnit> m_Neuron;

    std::string m_FilePath;
};

class Swc {
public:
    Swc(std::string filePath)
            : m_FilePath(filePath) {

    }

    void ReadFromFile() {
        std::ifstream infile;
        infile.open(m_FilePath);
        if (!infile.is_open()) {
            return;
        }

        std::string rowContent;
        while (std::getline(infile, rowContent)) {
            auto splitResult = string_split(rowContent, ' ');

            if (rowContent.empty() || rowContent[0] == '#') {
                rowContent.clear();
                continue;
            }

            if (splitResult.size() < 7) {
                continue;
            }
            NeuronUnit unit;
            unit.n = std::stoi(splitResult[0]);
            unit.type = std::stoi(splitResult[1]);
            unit.x = std::stof(splitResult[2]);
            unit.y = std::stof(splitResult[3]);
            unit.z = std::stof(splitResult[4]);
            unit.radius = std::stof(splitResult[5]);
            unit.parent = std::stoi(splitResult[6]);

            m_Neuron.push_back(unit);

            rowContent.clear();
        }
    }

    bool WriteToFile() {
        std::ofstream outfile;
        outfile.open(m_FilePath);
        if (!outfile.is_open()) {
            return false;
        }

        outfile << "#name\n"
                << "#Generated by SwcManagerClient\n"
                << "##n,type,x,y,z,radius,parent\n";

        for (auto &neuron: m_Neuron) {
            outfile << std::to_string(neuron.n) + " " +
                       std::to_string(neuron.type) + " " +
                       std::to_string(neuron.x) + " " +
                       std::to_string(neuron.y) + " " +
                       std::to_string(neuron.z) + " " +
                       std::to_string(neuron.radius) + " " +
                       std::to_string(neuron.parent)<<"\n";
        }
        outfile.close();
        return true;
    }

    std::vector<NeuronUnit> &getNeuron() {
        return m_Neuron;
    }

    void setNeuron(std::vector<NeuronUnit>& neuron) {
        m_Neuron = neuron;
    }

private:
    std::vector<NeuronUnit> m_Neuron;

    std::string m_FilePath;
};
