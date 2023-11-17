cmake_minimum_required(VERSION 3.26)
project(SwcManager)

set(CMAKE_CXX_STANDARD 20)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_PREFIX_PATH "C:/Qt/6.6.0/msvc2019_64")

find_package(Qt6 COMPONENTS
        Core
        Gui
        Widgets
        REQUIRED)

include_directories(
        ThirdParty/upb_x64-windows/include
        ThirdParty/abseil_x64-windows/include
        ThirdParty/c-ares_x64-windows/include
        ThirdParty/grpc_x64-windows/include
        ThirdParty/openssl_x64-windows/include
        ThirdParty/protobuf_x64-windows/include
        ThirdParty/re2_x64-windows/include
        ThirdParty/zlib_x64-windows/include
        ../Generated
)

if (CMAKE_BUILD_TYPE STREQUAL "Debug")
    link_directories(
            ThirdParty/upb_x64-windows/debug/lib
            ThirdParty/abseil_x64-windows/debug/lib
            ThirdParty/c-ares_x64-windows/debug/lib
            ThirdParty/grpc_x64-windows/debug/lib
            ThirdParty/openssl_x64-windows/debug/lib
            ThirdParty/protobuf_x64-windows/debug/lib
            ThirdParty/re2_x64-windows/debug/lib
            ThirdParty/zlib_x64-windows/debug/lib
    )
else ()

endif ()

set(PROJECT_HEADERS

)

set(PROJECT_SOURCES
        main.cpp

        ../Generated/Message/Message.grpc.pb.cc
        ../Generated/Message/Message.pb.cc
        ../Generated/Message/Request.grpc.pb.cc
        ../Generated/Message/Request.pb.cc
        ../Generated/Message/Response.grpc.pb.cc
        ../Generated/Message/Response.pb.cc
        ../Generated/Service/Service.grpc.pb.cc
        ../Generated/Service/Service.pb.cc

)

add_executable(SwcManager ${PROJECT_HEADERS} ${PROJECT_SOURCES})

if (CMAKE_BUILD_TYPE STREQUAL "Debug")
    target_link_libraries(SwcManager
            Qt::Core
            Qt::Gui
            Qt::Widgets
            abseil_dll
            cares
            address_sorting gpr grpc grpc++
            libcrypto libssl
            libprotobufd
            re2
            descriptor_upb_proto upb upb_textformat upb_extension_registry upb_fastdecode upb_json upb_mini_table upb_textformat  upb_reflection upb_collections upb_utf8_range
            zlibd
    )
else ()

endif ()

if (WIN32 AND NOT DEFINED CMAKE_TOOLCHAIN_FILE)
    set(DEBUG_SUFFIX)
    if (MSVC AND CMAKE_BUILD_TYPE MATCHES "Debug")
        set(DEBUG_SUFFIX "d")
    endif ()
    set(QT_INSTALL_PATH "${CMAKE_PREFIX_PATH}")
    if (NOT EXISTS "${QT_INSTALL_PATH}/bin")
        set(QT_INSTALL_PATH "${QT_INSTALL_PATH}/..")
        if (NOT EXISTS "${QT_INSTALL_PATH}/bin")
            set(QT_INSTALL_PATH "${QT_INSTALL_PATH}/..")
        endif ()
    endif ()
    if (EXISTS "${QT_INSTALL_PATH}/plugins/platforms/qwindows${DEBUG_SUFFIX}.dll")
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E make_directory
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>/plugins/platforms/")
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy
                "${QT_INSTALL_PATH}/plugins/platforms/qwindows${DEBUG_SUFFIX}.dll"
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>/plugins/platforms/")
    endif ()
    foreach (QT_LIB Core Gui Widgets)
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy
                "${QT_INSTALL_PATH}/bin/Qt6${QT_LIB}${DEBUG_SUFFIX}.dll"
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>")
    endforeach (QT_LIB)
endif ()
