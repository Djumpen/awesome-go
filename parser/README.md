# Awesome Go

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome)

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Files](#files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Functional](#functional)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Geographic](#geographic)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot-internet-of-things)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Miscellaneous](#miscellaneous)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [Validation](#validation)
    - [Version Control](#version-control)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
        - [Routers](#routers)
    - [Windows](#windows)
    - [XML](#xml)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Server Applications](#server-applications)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) - EasyMidi is a simple and reliable library for working with standard midi file (SMF). **⭐15**
* [flac](https://github.com/eaburns/flac) - Native Go FLAC decoder. **⭐79**
* [flac](https://github.com/mewkiz/flac) - Native Go FLAC decoder. **⭐78**
* [gaad](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser. **⭐50**
* [go-sox](https://github.com/krig/go-sox) - libsox bindings for go. **⭐81**
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) - libmediainfo bindings for go. **⭐19**
* [gosamplerate](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go. **⭐8**
* [id3v2](https://github.com/bogem/id3v2) - Fast and stable ID3 parsing and writing library for Go. **⭐80**
* [malgo](https://github.com/gen2brain/malgo) - Mini audio library. **⭐41**
* [minimp3](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library. **⭐15**
* [mix](https://github.com/go-mix/mix) - Sequence-based Go-native audio mixer for music apps. **⭐87**
* [mp3](https://github.com/tcolgate/mp3) - Native Go MP3 decoder. **⭐79**
* [music-theory](https://github.com/go-music-theory/music-theory) - Music theory models in Go. **⭐219**
* [PortAudio](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library. **⭐222**
* [portmidi](https://github.com/rakyll/portmidi) - Go bindings for PortMidi. **⭐178**
* [taglib](https://github.com/wtolson/go-taglib) - Go bindings for taglib. **⭐64**
* [vorbis](https://github.com/mccoyst/vorbis) - "Native" Go Vorbis decoder (uses CGO, but has no dependencies). **⭐22**
* [waveform](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams. **⭐221**

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. **⭐1596**
* [branca](https://github.com/hako/branca) - Golang implementation of Branca Tokens. **⭐33**
* [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, ABAC. **⭐3218**
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) - provides parser of cookies.txt file format. **⭐1**
* [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth) - AWS (Amazon Web Services) request signing library. **⭐215**
* [go-jose](https://github.com/square/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. **⭐975**
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) - Standalone, specification-compliant,  OAuth2 server written in Golang. **⭐1008**
* [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. **⭐929**
* [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. **⭐783**
* [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. **⭐1925**
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middleware. **⭐157**
* [jwt](https://github.com/robbert229/jwt) - Clean and easy to use implementation of JSON Web Tokens (JWT). **⭐50**
* [jwt](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library. **⭐14**
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options. **⭐133**
* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT). **⭐4484**
* [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. **⭐653**
* [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. **⭐1906**
* [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. **⭐1481**
* [paseto](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO) **⭐156**
* [permissions2](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. **⭐285**
* [rbac](https://github.com/zpatrick/rbac) - Minimalistic RBAC package for Go applications. **⭐16**
* [securecookie](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding. **⭐24**
* [session](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE). **⭐74**
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) - Go session management using the SessionGate Redis module. **⭐5**
* [sessions](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers. **⭐36**
* [signedvalue](https://github.com/sashka/signedvalue) - Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. **⭐5**
* [yubigo](https://github.com/GeertJohan/yubigo) - Yubikey client package that provides a simple API to integrate the Yubico Yubikey into a go application. **⭐95**

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [argparse](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module. **⭐58**
* [argv](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax. **⭐10**
* [cli](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags. **⭐405**
* [cli](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go. **⭐37**
* [cli-init](https://github.com/tcnksm/gcli) - The easy way to start building Golang command line applications. **⭐845**
* [climax](http://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command.
* [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. **⭐9512**
* [commandeer](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. **⭐66**
* [complete](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion. **⭐517**
* [docopt.go](https://github.com/docopt/docopt.go) - Command-line arguments parser that will make you smile.
* [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. **⭐4492**
* [env](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs. **⭐34**
* [flag](https://github.com/cosiner/flag) - Simple but powerful command line option parsing library for Go supporting subcommand. **⭐85**
* [flaggy](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support. **⭐389**
* [flagvar](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package. **⭐23**
* [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. **⭐538**
* [go-commander](https://github.com/yitsushi/go-commander) - Go library to simplify CLI workflow **⭐7**
* [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. **⭐1203**
* [gocmd](https://github.com/devfacet/gocmd) - Go library for building command line applications. **⭐23**
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection.
* [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands. **⭐2119**
* [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. **⭐490**
* [mitchellh/cli](https://github.com/mitchellh/cli) - Go library for implementing command-line interfaces. **⭐878**
* [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation.
* [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. **⭐537**
* [readline](https://github.com/chzyer/readline) - Pure golang implementation that provides most features in GNU-Readline under MIT license. **⭐1245**
* [sflags](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. **⭐64**
* [strumt](https://github.com/antham/strumt) - Library to create prompt chain. **⭐20**
* [ukautz/clif](https://github.com/ukautz/clif) - Small command line interface framework. **⭐89**
* [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). **⭐9286**
* [wlog](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency. **⭐26**
* [wmenu](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompts users to make choices. **⭐59**

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that supports fmt.Printf/Sprintf. **⭐378**
* [cfmt](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes. **⭐45**
* [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. **⭐278**
* [color](https://github.com/fatih/color) - Versatile package for colored terminal output. **⭐2749**
* [colourize](https://github.com/TreyBastian/colourize) - Go library for ANSI colour text in terminals. **⭐13**
* [ctc](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method **⭐4**
* [go-ataman](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals. **⭐7**
* [go-colorable](https://github.com/mattn/go-colorable) - Colorable writer for windows. **⭐310**
* [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals. **⭐179**
* [go-isatty](https://github.com/mattn/go-isatty) - isatty for golang. **⭐280**
* [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). **⭐1875**
* [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. **⭐3655**
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
* [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. **⭐31**
* [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. **⭐355**
* [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. **⭐420**
* [simpletable](https://github.com/alexeyco/simpletable) - Simple tables in terminal with Go. **⭐103**
* [tabular](https://github.com/InVisionApp/tabular) - Print ASCII tables from command line utilities without the need to pass large sets of data to the API. **⭐14**
* [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. **⭐3076**
* [termtables](https://github.com/apcera/termtables) - Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output. **⭐194**
* [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). **⭐8019**
* [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in realtime. **⭐690**
* [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. **⭐1201**
* [uitable](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data. **⭐451**

## Configuration

*Libraries for configuration parsing.*

* [config](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing. **⭐178**
* [configure](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables. **⭐40**
* [confita](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct. **⭐180**
* [conflate](https://github.com/miracl/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. **⭐15**
* [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). **⭐621**
* [envcfg](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs. **⭐88**
* [envconf](https://github.com/ian-kent/envconf) - Configuration from environment. **⭐7**
* [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables. **⭐127**
* [envh](https://github.com/antham/envh) - Helpers to manage environment variables. **⭐89**
* [gcfg](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections. **⭐100**
* [go-up](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic. **⭐15**
* [goConfig](https://github.com/crgimenes/goConfig) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. **⭐83**
* [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). **⭐1539**
* [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy. **⭐55**
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gookit/config](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. **⭐24**
* [hjson](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. **⭐150**
* [ingo](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file. **⭐21**
* [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. **⭐1181**
* [joshbetz/config](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. **⭐190**
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. **⭐1931**
* [mini](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files. **⭐19**
* [sprbox](https://github.com/oblq/sprbox) - Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). **⭐3**
* [store](https://github.com/tucnak/store) - Lightweight configuration manager for Go. **⭐236**
* [viper](https://github.com/spf13/viper) - Go configuration with fangs. **⭐6781**
* [xdg](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). **⭐14**

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. **⭐16207**
* [duci](https://github.com/duck8823/duci) - A simple ci server no needs domain specific languages. **⭐14**
* [gomason](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace. **⭐12**
* [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system. **⭐521**
* [overalls](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls. **⭐94**
* [roveralls](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool. **⭐10**

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor. **⭐408**
* [go-libsass](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project. **⭐107**

## Data Structures

*Generic datastructures and algorithms in Go.*
* [algorithms](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study. **⭐97**
* [binpacker](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream. **⭐99**
* [bit](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions. **⭐24**
* [bitset](https://github.com/willf/bitset) - Go package implementing bitsets. **⭐403**
* [bloom](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. **⭐119**
* [bloom](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation. **⭐10**
* [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. **⭐1031**
* [concurrent-writer](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`. **⭐11**
* [conjungo](https://github.com/InVisionApp/conjungo) - A small, powerful and flexible merge library. **⭐45**
* [count-min-log](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). **⭐39**
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. **⭐389**
* [deque](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue). **⭐12**
* [encoding](https://github.com/zhenjl/encoding) - Integer Compression Libraries for Go. **⭐88**
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree. **⭐55**
* [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. **⭐4520**
* [go-ef](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding. **⭐8**
* [go-geoindex](https://github.com/hailocab/go-geoindex) - In-memory geo index. **⭐287**
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches. **⭐16**
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding. **⭐88**
* [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. **⭐4764**
* [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. **⭐888**
* [goset](https://github.com/zoumo/goset) - A useful Set collection implementation for Go. **⭐10**
* [goskiplist](https://github.com/ryszard/goskiplist) - Skip list implementation in Go. **⭐172**
* [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. **⭐665**
* [hilbert](https://github.com/google/hilbert) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. **⭐160**
* [hyperloglog](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. **⭐619**
* [levenshtein](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. **⭐21**
* [levenshtein](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go. **⭐33**
* [mafsa](https://github.com/smartystreets/mafsa) - MA-FSA implementation with Minimal Perfect Hashing. **⭐269**
* [merkletree](https://github.com/cbergoon/merkletree) - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. **⭐101**
* [mspm](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval. **⭐3**
* [pipeline](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out. **⭐6**
* [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. **⭐483**
* [set](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap. **⭐5**
* [skiplist](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation. **⭐62**
* [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go. **⭐57**
* [trie](https://github.com/derekparker/trie) - Trie implementation in Go. **⭐327**
* [ttlcache](https://github.com/diegobernardes/ttlcache) - In-memory LRU string-interface{} map with expiration for golang. **⭐73**
* [willf/bloom](https://github.com/willf/bloom) - Go package implementing Bloom filters. **⭐546**

## Database

*Databases implemented in Go.*

* [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. **⭐4586**
* [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. **⭐1699**
* [bolt](https://github.com/boltdb/bolt) - Low-level key/value database for Go. **⭐9143**
* [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. **⭐2135**
* [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. **⭐692**
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration. **⭐23**
* [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. **⭐14712**
* [couchcache](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server. **⭐38**
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. **⭐357**
* [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. **⭐6926**
* [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. **⭐653**
* [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. **⭐494**
* [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB. **⭐29**
* [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. **⭐566**
* [go-cache](https://github.com/pmylund/go-cache) - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. **⭐2091**
* [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. **⭐2586**
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) - Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. **⭐6**
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. **⭐6862**
* [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. **⭐14841**
* [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. **⭐2747**
* [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB. **⭐338**
* [moss](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go. **⭐649**
* [piladb](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures. **⭐165**
* [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. **⭐20208**
* [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. **⭐4100**
* [Scribble](https://github.com/nanobox-io/golang-scribble) - Tiny flat file JSON store. **⭐17**
* [slowpoke](https://github.com/recoilme/slowpoke) - Key-value store with persistence. **⭐65**
* [tempdb](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items. **⭐13**
* [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. **⭐15910**
* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. **⭐2242**
* [Vasto](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. **⭐99**

*Database schema migration.*

* [darwin](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go. **⭐71**
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library. **⭐17**
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg. **⭐11**
* [gondolier](https://github.com/emvicom/gondolier) - Gondolier is a library to auto migrate database schemas using structs. **⭐20**
* [goose](https://github.com/steinbacher/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. **⭐86**
* [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. **⭐211**
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. **⭐1118**
* [pravasan](https://github.com/pravasan/pravasan) - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. **⭐22**
* [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. **⭐1140**

*Database tools.*

* [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. **⭐178**
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - Collects small insterts and sends big requests to ClickHouse servers. **⭐87**
* [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. **⭐1354**
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically. **⭐1800**
* [kingshard](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang. **⭐3882**
* [myreplication](https://github.com/2tvenom/myreplication) - MySql binary log replication listener. Supports statement and row based replication. **⭐123**
* [orchestrator](https://github.com/github/orchestrator) - MySQL replication topology manager & visualizer. **⭐2256**
* [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. **⭐5381**
* [prep](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code. **⭐20**
* [pREST](https://github.com/nuveo/prest) - Serve a RESTful API from any PostgreSQL database. **⭐1879**
* [rwdb](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup. **⭐9**
* [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. **⭐6900**

*SQL query builder, libraries for building and using SQL.*

* [dat](https://github.com/mgutz/dat) - Go Postgres Data Access Toolkit. **⭐544**
* [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease. **⭐383**
* [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. **⭐460**
* [godbal](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. **⭐48**
* [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. **⭐462**
* [igor](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. **⭐70**
* [ormlite](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases. **⭐0**
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities. **⭐273**
* [scaneo](https://github.com/variadico/scaneo) - Generate Go code to convert database rows into arbitrary structs. **⭐136**
* [sqrl](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance. **⭐132**
* [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. **⭐1769**
* [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. **⭐1806**

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [avatica](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql. **⭐19**
    * [bgc](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go. **⭐9**
    * [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go. **⭐89**
    * [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql. **⭐84**
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. **⭐837**
    * [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql. **⭐335**
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. **⭐6435**
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. **⭐2855**
    * [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). **⭐77**
    * [goracle](https://github.com/go-goracle/goracle) - Oracle driver for Go, using the ODPI-C driver **⭐139**
    * [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. **⭐1541**
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. **⭐4358**

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language. **⭐285**
    * [arangolite](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB. **⭐65**
    * [asc](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go. **⭐4**
    * [cachego](https://github.com/fabiorphp/cachego) - Golang Cache component for multiple drivers. **⭐98**
    * [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. **⭐11908**
    * [dsc](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files. **⭐6**
    * [dynago](https://github.com/underarmour/dynago) - Dynago is a principle of least surprise client for DynamoDB. **⭐61**
    * [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go. **⭐283**
    * [go-couchdb](https://github.com/fjl/go-couchdb) - Yet another CouchDB HTTP API wrapper for Go. **⭐50**
    * [gocb](https://github.com/couchbase/gocb) - Official Couchbase Go SDK. **⭐272**
    * [gocql](http://gocql.github.io) - Go language driver for Apache Cassandra.
    * [godscache](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. **⭐3**
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language. **⭐976**
    * [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. **⭐1388**
    * [goriak](https://github.com/zegl/goriak) - Go language driver for Riak KV. **⭐23**
    * [mgo](https://github.com/globalsign/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms **⭐1231**
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. **⭐1139**
    * [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang. **⭐24**
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang. **⭐70**
    * [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang. **⭐344**
    * [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. **⭐5153**
    * [redis](https://github.com/go-redis/redis) - Redis client for Golang. **⭐4475**
    * [redis](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services. **⭐227**
    * [xredis](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client. **⭐9**

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. **⭐4727**
    * [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. **⭐3135**
    * [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. **⭐235**
    * [elastigo](https://github.com/mattbaird/elastigo) - Elasticsearch client library. **⭐923**
    * [goes](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch. **⭐24**
    * [riot](https://github.com/go-ego/riot) - Go Open Source, Distributed, Simple and efficient Search Engine **⭐4075**
    * [skizze](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage. **⭐55**

## Date and Time

*Libraries for working with dates and times.*

* [carbon](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library. **⭐265**
* [date](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. **⭐21**
* [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. **⭐740**
* [durafmt](https://github.com/hako/durafmt) - Time duration formatting library for Go. **⭐201**
* [feiertage](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... **⭐15**
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang). **⭐38**
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location. **⭐8**
* [goweek](https://github.com/grsmv/goweek) - Library for working with week entity in golang. **⭐17**
* [Kair](https://github.com/GuilhermeCaruso/Kair) - Date and Time - Golang Formatting Library. **⭐5**
* [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. **⭐1855**
* [NullTime](https://github.com/kirillDanshin/nulltime) - Nullable `time.Time`. **⭐7**
* [strftime](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter. **⭐5**
* [timespan](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration. **⭐57**
* [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package. **⭐150**
* [tuesday](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function. **⭐6**

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
* [consistent](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads. **⭐136**
* [digota](https://github.com/digota/digota) - grpc ecommerce microservice. **⭐249**
* [doublejump](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash. **⭐3**
* [drmaa](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard. **⭐21**
* [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. **⭐1435**
* [flowgraph](https://github.com/vectaport/flowgraph) - flow-based programming package. **⭐3**
* [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. **⭐1652**
* [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. **⭐2228**
* [go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service. **⭐386**
* [go-jump](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function. **⭐211**
* [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. **⭐11831**
* [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load. **⭐494**
* [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. **⭐6863**
* [heimdall](https://github.com/gojektech/heimdall) - An enchanced http client with retry and hystrix capabilities. **⭐603**
* [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. **⭐860**
* [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. **⭐6390**
* [jsonrpc](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0. **⭐85**
* [jsonrpc](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation. **⭐65**
* [KrakenD](https://github.com/devopsfaith/krakend) - Ultra performant API Gateway framework with middlewares. **⭐1033**
* [micro](https://github.com/micro/micro) - Pluggable microservice toolkit and distributed systems platform. **⭐5033**
* [NATS](https://github.com/nats-io/gnatsd) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. **⭐4840**
* [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. **⭐2231**
* [raft](https://github.com/coreos/etcd/tree/master/raft) - Go implementation of the Raft consensus protocol, by CoreOS.
* [redis-lock](https://github.com/bsm/redis-lock) - Simplified distributed locking implementation using Redis. **⭐81**
* [ringpop-go](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications. **⭐505**
* [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. **⭐2942**
* [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). **⭐279**
* [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. **⭐2479**
* [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. **⭐2482**
    * [dht](https://godoc.org/github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation.
    * [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client. **⭐339**

## Email

*Libraries and tools that implement email creation and sending.*

* [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.
* [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails. **⭐135**
* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. **⭐952**
* [go-dkim](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email. **⭐45**
* [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. **⭐547**
* [go-message](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages. **⭐70**
* [Gomail](https://github.com/go-gomail/gomail/) - Gomail is a very simple and powerful package to send emails. **⭐1952**
* [Hectane](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API. **⭐159**
* [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. **⭐1368**
* [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. **⭐4041**
* [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email. **⭐420**
* [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine. **⭐49**

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [agora](https://github.com/PuerkitoBio/agora) - Dynamically typed, embeddable programming language in Go. **⭐312**
* [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. **⭐783**
* [binder](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). **⭐22**
* [expr](https://github.com/antonmedv/expr) - an engine that can evaluate expressions. **⭐215**
* [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go. **⭐407**
* [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go. **⭐608**
* [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. **⭐1417**
* [go-php](https://github.com/deuill/go-php) - PHP bindings for Go. **⭐539**
* [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. **⭐761**
* [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API. **⭐411**
* [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. **⭐2419**
* [ngaro](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro. **⭐15**
* [otto](https://github.com/robertkrimen/otto) - JavaScript interpreter written in Go. **⭐4187**
* [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go. **⭐24**

## Files

*Libraries for  handling files and file systems.*

* [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. **⭐1760**
* [go-csv-tag](https://github.com/artonge/go-csv-tag) - Load csv file using tag. **⭐34**
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans. **⭐6**
* [go-gtfs](https://github.com/artonge/go-gtfs) - Load gtfs files in go. **⭐9**
* [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal. **⭐403**
* [pdfcpu](https://github.com/hhrutter/pdfcpu) - PDF processor. **⭐704**
* [skywalker](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease. **⭐38**
* [tarfs](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. **⭐25**

## Financial

*Packages for accounting and finance.*

* [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang. **⭐415**
* [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. **⭐1176**
* [go-finance](https://github.com/FlashBoys/go-finance) - Comprehensive financial markets data in Go. **⭐540**
* [go-finance](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. **⭐26**
* [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. **⭐493**
* [ofxgo](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client). **⭐48**
* [techan](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies. **⭐78**
* [transaction](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode. **⭐31**
* [vat](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates. **⭐51**

## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind) - Bind form data to any Go values. **⭐20**
* [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct. **⭐717**
* [conform](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. **⭐148**
* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. **⭐302**
* [formam](https://github.com/monoculum/formam) - decode form's values into a struct. **⭐109**
* [forms](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. **⭐94**
* [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. **⭐333**
* [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. **⭐883**

## Functional

*Packages to support functional programming in Go.*

* [fpGo](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang **⭐56**
* [fuego](https://github.com/seborama/fuego) - Functional Experiment in Go **⭐0**
* [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. **⭐979**

## Game Development

*Awesome game development libraries.*

* [Azul3D](https://github.com/azul3d/engine) - 3D game engine written in Go. **⭐392**
* [Ebiten](https://github.com/hajimehoshi/ebiten) - dead simple 2D game library in Go. **⭐1175**
* [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. **⭐891**
* [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. **⭐470**
* [GarageEngine](https://github.com/vova616/GarageEngine) - 2d game engine written in Go working on OpenGL. **⭐304**
* [glop](https://github.com/runningwild/glop) - Glop (Game Library Of Power) is a fairly simple cross-platform game library. **⭐76**
* [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm. **⭐289**
* [go-collada](https://github.com/GlenKelley/go-collada) - Go package for working with the Collada file format. **⭐12**
* [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). **⭐984**
* [go3d](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go. **⭐142**
* [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. **⭐963**
* [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping **⭐837**
* [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. **⭐2491**
* [nano](https://github.com/lonnng/nano) - Lightweight, facility, high performance golang based game server framework **⭐736**
* [Oak](https://github.com/oakmound/oak) - Pure Go game engine. **⭐540**
* [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. **⭐79**
* [Pixel](https://github.com/faiface/pixel) - Hand-crafted 2D game library in Go. **⭐1802**
* [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. **⭐294**
* [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. **⭐931**

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [efaceconv](https://github.com/t0pep0/efaceconv) - Code generation tool for high performance conversion from interface{} to immutable type without allocations. **⭐35**
* [gen](https://github.com/clipperhouse/gen) - Code generation tool for ‘generics’-like functionality. **⭐940**
* [go-enum](https://github.com/abice/go-enum) - Code generation for enums from code comments. **⭐61**
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. **⭐1554**
* [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types. **⭐628**
* [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. **⭐148**
* [interfaces](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions. **⭐144**
* [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. **⭐977**
* [pkgreflect](https://github.com/ungerik/pkgreflect) - Go preprocessor for package scoped reflection. **⭐72**

## Geographic

*Geographic tools and servers*

* [geocache](https://github.com/melihmucuk/geocache) - In-memory cache that is suitable for geolocation based applications. **⭐93**
* [geoserver](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. **⭐14**
* [gismanager](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver **⭐0**
* [osm](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs. **⭐28**
* [pbf](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder. **⭐9**
* [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go.
* [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. **⭐5007**

## Go Compilers

*Tools for compiling Go to other languages.*

* [c4go](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code. **⭐43**
* [f4go](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code. **⭐7**
* [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. **⭐7569**
* [llgo](https://github.com/go-llvm/llgo) - LLVM-based compiler for Go. **⭐931**
* [tardisgo](https://github.com/tardisgo/tardisgo) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. **⭐381**

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) - A high-performance goroutine pool for golang. **⭐935**
* [artifex](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching **⭐1**
* [async](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic. **⭐11**
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang. **⭐20**
* [go-floc](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease. **⭐159**
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order. **⭐86**
* [go-trylock](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang. **⭐4**
* [GoSlaves](https://github.com/themester/GoSlaves) - Simple and Asynchronous Goroutine pool library. **⭐48**
* [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. **⭐2068**
* [grpool](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool. **⭐428**
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel. **⭐21**
* [pool](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. **⭐403**
* [semaphore](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. **⭐36**
* [semaphore](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). **⭐45**
* [stl](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. **⭐5**
* [threadpool](https://github.com/shettyh/threadpool) - Golang threadpool implementation. **⭐6**
* [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. **⭐963**
* [worker-pool](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool. **⭐20**
* [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. **⭐52**

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. **⭐2545**
* [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows. **⭐208**
* [go-astilectron](https://github.com/asticode/go-astilectron) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). **⭐2069**
* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. **⭐1204**
* [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. **⭐549**
* [gowd](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. **⭐151**
* [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). **⭐4782**
* [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. **⭐6133**
* [walk](https://github.com/lxn/walk) - Windows application library kit for Go. **⭐3001**
* [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). **⭐3132**

*Interaction*

* [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go. **⭐459**
* [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. **⭐3566**
* [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. **⭐553**
* [trayhost](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. **⭐136**


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. **⭐1850**
* [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. **⭐614**
* [cameron](https://github.com/aofei/cameron) - An avatar generator for Go. **⭐9**
* [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. **⭐968**
* [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. **⭐1551**
* [gift](https://github.com/disintegration/gift) - Package of image processing filters. **⭐1133**
* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library. **⭐79**
* [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library. **⭐47**
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go. **⭐269**
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV. **⭐992**
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go. **⭐24**
* [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. **⭐1695**
* [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package. **⭐147**
* [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars. **⭐266**
* [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII. **⭐203**
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. **⭐810**
* [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. **⭐2039**
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. **⭐1978**
* [img](https://github.com/hawx/img) - Selection of image manipulation tools. **⭐125**
* [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. **⭐2299**
* [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). **⭐34**
* [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go. **⭐336**
* [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos. **⭐4**
* [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. **⭐901**
* [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. **⭐1664**
* [resize](https://github.com/nfnt/resize) - Image resizing for Go with common interpolation methods. **⭐1932**
* [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD. **⭐155**
* [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. **⭐1180**
* [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography. **⭐10**
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. **⭐1172**
* [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder. **⭐19**

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT. **⭐122**
* [devices](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io. **⭐211**
* [eywa](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices. **⭐22**
* [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. **⭐825**
* [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. **⭐722**
* [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things. **⭐4866**
* [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device. **⭐23**
* [mainflux](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server. **⭐222**
* [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.
* [sensorbee](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT. **⭐158**

## Logging

*Libraries for generating and working with log files.*

* [distillog](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels). **⭐14**
* [glg](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go. **⭐23**
* [glog](https://github.com/golang/glog) - Leveled execution logs for Go. **⭐2028**
* [go-cronowriter](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog. **⭐15**
* [go-log](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers. **⭐6**
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers. **⭐19**
* [go-log](https://github.com/ian-kent/go-log) - Log4j implementation in Go. **⭐31**
* [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers. **⭐199**
* [gologger](https://github.com/sadlil/gologger) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. **⭐38**
* [gomol](https://github.com/aphistic/gomol) - Multiple-output, structured logging for Go with extensible logging outputs. **⭐12**
* [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.
* [journald](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging. **⭐14**
* [log](https://github.com/apex/log) - Structured logging package for Go. **⭐623**
* [log](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go. **⭐255**
* [log](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation. **⭐19**
* [log-voyage](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang. **⭐82**
* [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. **⭐817**
* [logdump](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging. **⭐7**
* [logex](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib. **⭐33**
* [logger](https://github.com/azer/logger) - Minimalistic logging library for Go. **⭐129**
* [logmatic](https://github.com/borderstech/logmatic) - Colorized logger for Golang with dynamic log level configuration **⭐1**
* [logo](https://github.com/mbndr/logo) - Golang logger to different configurable writers. **⭐4**
* [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. **⭐8954**
* [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). **⭐24**
* [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger. **⭐228**
* [logxi](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy. **⭐327**
* [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. **⭐1047**
* [mlog](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. **⭐16**
* [onelog](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. **⭐291**
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). **⭐97**
* [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. **⭐1217**
* [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. **⭐2700**
* [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. **⭐44**
* [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. **⭐1243**
* [xlog](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. **⭐3**
* [xlog](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching. **⭐128**
* [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. **⭐5513**
* [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. **⭐1499**

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. **⭐578**
* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. **⭐619**
* [eaopt](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library. **⭐560**
* [fonet](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go. **⭐21**
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms. **⭐14**
* [go-deep](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go. **⭐178**
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library. **⭐97**
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang. **⭐162**
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. **⭐55**
* [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go. **⭐296**
* [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods. **⭐20**
* [goga](https://github.com/tomcraven/goga) - Genetic algorithm library for Go. **⭐69**
* [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. **⭐6085**
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go. **⭐36**
* [GoMind](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go. **⭐5**
* [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. **⭐925**
* [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go. **⭐125**
* [gorgonia](https://github.com/chewxy/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. **⭐2248**
* [goscore](https://github.com/asafschers/goscore) - Go Scoring API for PMML. **⭐21**
* [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. **⭐660**
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14. **⭐58**
* [mlgo](https://github.com/NullHypothesis/mlgo) - This project aims to provide minimalistic machine learning algorithms in Go. **⭐4**
* [neat](https://github.com/jinyeom/neat) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). **⭐47**
* [neural-go](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation. **⭐58**
* [probab](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go. **⭐9**
* [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine. **⭐220**
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go. **⭐118**
* [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. **⭐990**
* [Varis](https://github.com/Xamber/Varis) - Golang Neural Network. **⭐17**

## Messaging

*Libraries that implement messaging systems.*

* [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. **⭐1873**
* [Benthos](https://github.com/Jeffail/benthos) - A message streaming bridge between a range of protocols. **⭐1618**
* [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. **⭐3133**
* [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. **⭐286**
* [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. **⭐53**
* [emitter](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. **⭐251**
* [event](https://github.com/agoalofalife/event) - Implementation of the pattern observer. **⭐17**
* [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. **⭐445**
* [gaurun-client](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go. **⭐7**
* [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io). **⭐293**
* [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec. **⭐38**
* [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. **⭐1221**
* [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework.
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service. **⭐10**
* [Gollum](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. **⭐633**
* [golongpoll](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple. **⭐386**
* [goose](https://github.com/ian-kent/goose) - Server Sent Events in Go. **⭐34**
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. **⭐1749**
* [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). **⭐3049**
* [guble](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. **⭐133**
* [hub](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. **⭐12**
* [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. **⭐2681**
* [mangos](https://github.com/go-mangos/mangos) - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. **⭐1478**
* [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. **⭐1189**
* [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). **⭐377**
* [messagebus](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. **⭐25**
* [NATS Go Client](https://github.com/nats-io/nats) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel. **⭐39**
* [oplog](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs. **⭐91**
* [pubsub](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go. **⭐224**
* [rabbus](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues. **⭐46**
* [rabtap](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app. **⭐33**
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue. **⭐48**
* [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. **⭐3610**
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. **⭐1029**
* [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). **⭐698**

## Miscellaneous

*These libraries were placed here because none of the other categories seemed to fit.*

* [alice](https://github.com/magic003/alice) - Additive dependency injection container for Golang. **⭐29**
* [anagent](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. **⭐8**
* [antch](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework. **⭐110**
* [archiver](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives. **⭐1690**
* [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields. **⭐21**
* [avgRating](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation. **⭐8**
* [banner](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications. **⭐190**
* [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. **⭐423**
* [battery](https://github.com/distatus/battery) - Cross-platform, normalized battery information library. **⭐107**
* [bitio](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go. **⭐68**
* [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](http://browscap.org/). **⭐27**
* [captcha](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation. **⭐25**
* [conv](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types. **⭐330**
* [datacounter](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter. **⭐25**
* [errors](https://github.com/pkg/errors) - Package that provides simple error handling primitives. **⭐3645**
* [ffmt](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans. **⭐85**
* [ghorg](https://github.com/gabrie30/ghorg) - Clone all repos from a GitHub org into a single directory. **⭐13**
* [go-chat-bot](https://github.com/go-chat-bot/bot) - IRC, Slack & Telegram bot written in Go. **⭐371**
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. **⭐554**
* [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. **⭐516**
* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.
* [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. **⭐704**
* [go-sarah](https://github.com/oklahomer/go-sarah) - Framework to build bot for desired chat services including LINE, Slack, Gitter and more. **⭐109**
* [go-unarr](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives. **⭐44**
* [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. **⭐190**
* [goid](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs. **⭐13**
* [gommit](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns **⭐55**
* [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). **⭐3068**
* [gosh](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method. **⭐11**
* [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. **⭐1174**
* [gountries](https://github.com/pariz/gountries) - Package that exposes country and subdivision data. **⭐182**
* [hanu](https://github.com/sbstjn/hanu) - Framework for writing Slack bots. **⭐76**
* [health](https://github.com/dimiro1/health) - Easy to use, extensible health check library. **⭐328**
* [healthcheck](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services. **⭐54**
* [hostutils](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list. **⭐5**
* [indigo](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58. **⭐42**
* [jobs](https://github.com/albrow/jobs) - Persistent and flexible background jobs library. **⭐432**
* [lk](https://github.com/hyperboloide/lk) - A simple licensing library for golang. **⭐77**
* [margelet](https://github.com/zhulik/margelet) - Framework for building Telegram bots. **⭐49**
* [morse](https://github.com/alwindoss/morse) - Library to convert to and from morse code. **⭐38**
* [pdfgen](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests. **⭐17**
* [persian](https://github.com/mavihq/persian) - Some utilities for Persian language in go. **⭐22**
* [sandid](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID. **⭐3**
* [secdl](https://github.com/xor-gate/secdl) - Lighttpd ModSecDownload algorithm ported to go to secure download urls.
* [shellwords](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. **⭐3**
* [shortid](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs. **⭐348**
* [slacker](https://github.com/shomali11/slacker) - Easy to use framework to create Slack bots. **⭐219**
* [stats](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... **⭐108**
* [strutil](https://github.com/ozgio/strutil) - String utilities **⭐34**
* [turtle](https://github.com/hackebrot/turtle) - Emojis for Go. **⭐57**
* [url-shortener](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support. **⭐9**
* [uuid](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. **⭐6**
* [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. **⭐331**
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling.
* [werr](https://github.com/txgruppi/werr) - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. **⭐7**
* [wuid](https://github.com/edwingeng/wuid) - An extremely fast unique number generator, 10-135 times faster than UUID. **⭐215**
* [xdg](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go. **⭐7**
* [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber. **⭐36**
* [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. **⭐519**

## Natural Language Processing

*Libraries for working with human languages.*

* [getlang](https://github.com/rylans/getlang) - Fast natural language detection package. **⭐32**
* [go-eco](https://github.com/ThePaw/go-eco) - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models. **⭐4**
* [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text. **⭐764**
* [go-mystem](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer. **⭐19**
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work. **⭐76**
* [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. **⭐397**
* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm. **⭐50**
* [go-unidecode](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text. **⭐35**
* [go2vec](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings. **⭐30**
* [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. **⭐661**
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2. **⭐15**
* [gotokenizer](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) **⭐2**
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go. **⭐59**
* [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. **⭐618**
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. **⭐17**
* [kagome](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go. **⭐336**
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. **⭐10**
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. **⭐57**
* [nlp](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp. **⭐349**
* [nlp](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). **⭐168**
* [paicehusk](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm. **⭐24**
* [petrovich](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case. **⭐16**
* [porter](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. **⭐8**
* [porter2](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer. **⭐32**
* [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. **⭐1816**
* [RAKE.go](https://github.com/Obaied/RAKE.go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
* [segment](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/) **⭐43**
* [sentences](https://github.com/neurosnap/sentences) - Sentence tokenizer:  converts text into a list of sentences. **⭐239**
* [shamoji](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go. **⭐9**
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). **⭐21**
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. **⭐41**
* [textcat](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text. **⭐58**
* [whatlanggo](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). **⭐304**
* [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. **⭐891**

## Networking

*Libraries for working with various layers of the network.*

* [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. **⭐155**
* [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy. **⭐214**
* [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252). **⭐121**
* [cidranger](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go. **⭐303**
* [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. **⭐52**
* [dns](https://github.com/miekg/dns) - Go library for working with DNS. **⭐3181**
* [ether](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames. **⭐57**
* [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. **⭐161**
* [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. **⭐7389**
* [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. **⭐517**
* [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). **⭐399**
* [gNxI](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols. **⭐71**
* [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. **⭐574**
* [go-stun](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389). **⭐273**
* [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. **⭐1433**
* [golibwireshark](https://github.com/sunwxg/golibwireshark) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. **⭐12**
* [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. **⭐2288**
* [gopcap](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap. **⭐329**
* [goshark](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet. **⭐7**
* [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. **⭐366**
* [gotcp](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications. **⭐371**
* [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. **⭐412**
* [graval](https://github.com/koofr/graval) - Experimental FTP server framework. **⭐23**
* [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. **⭐3229**
* [jazigo](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices. **⭐102**
* [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. **⭐1749**
* [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. **⭐9012**
* [lhttp](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily. **⭐449**
* [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces. **⭐36**
* [llb](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. **⭐7**
* [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. **⭐454**
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). **⭐538**
* [packet](https://github.com/aerogo/packet) - Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. **⭐7**
* [peerdiscovery](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast **⭐301**
* [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it. **⭐37**
* [publicip](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress). **⭐16**
* [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. **⭐1922**
* [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface. **⭐238**
* [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. **⭐576**
* [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). **⭐867**
* [sslb](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. **⭐107**
* [stun](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol. **⭐154**
* [tcp_server](https://github.com/firstrow/tcp_server) - Go library for building tcp servers faster. **⭐229**
* [tspool](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server.  **⭐0**
* [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation. **⭐139**
* [water](https://github.com/songgao/water) - Simple TUN/TAP library. **⭐678**
* [winrm](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines. **⭐180**
* [xtcp](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. **⭐48**

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). **⭐536**
* [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. **⭐529**
* [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). **⭐127**
* [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. **⭐54**
* [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM. **⭐249**

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [go-pg](https://github.com/go-pg/pg) - PostgreSQL ORM with focus on PostgreSQL specific features and performance. **⭐2149**
* [go-queryset](https://github.com/jirfag/go-queryset) - 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. **⭐358**
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. **⭐81**
* [go-store](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go. **⭐90**
* [gomodel](https://github.com/cosiner/gomodel) - Lightweight, fast, orm-like library helps interactive with database. **⭐60**
* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. **⭐11163**
* [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. **⭐2901**
* [grimoire](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). **⭐30**
* [lore](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. **⭐4**
* [Marlow](https://github.com/dadleyy/marlow) - Generated ORM from project structs for compile time safety assurances. **⭐36**
* [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. **⭐511**
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM. **⭐516**
* [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. **⭐699**
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. **⭐1680**
* [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go. **⭐4027**
* [Zoom](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis. **⭐226**

## Package Management

*Official tooling for package management*

* [dep](https://github.com/golang/dep) - Go dependency tool. **⭐10974**
* [vgo](https://go.googlesource.com/vgo/) - Versioned Go.

*Unofficial libraries for package and dependency management.*

* [gigo](https://github.com/LyricalSecurity/gigo) - PIP-like dependency tool for golang, with support for private repositories and hashes. **⭐196**
* [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. **⭐7248**
* [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. **⭐5533**
* [gom](https://github.com/mattn/gom) - Go Manager - bundle for go. **⭐1339**
* [goop](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler. **⭐774**
* [gop](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH **⭐48**
* [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. **⭐2019**
* [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. **⭐4032**
* [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. **⭐1201**
* [gvt](https://github.com/FiloSottile/gvt) - `gvt` is a simple vendoring tool made for Go native vendoring (aka GO15VENDOREXPERIMENT), based on gb-vendor. **⭐754**
* [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git. **⭐213**
* [nut](https://github.com/jingweno/nut) - Vendor Go dependencies. **⭐247**
* [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments. **⭐111**

## Query Language

* [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. **⭐456**
* [graphql](https://github.com/tmc/graphql) - graphql parser + utilities. **⭐47**
* [graphql](https://github.com/sevki/graphql) - GraphQL implementation in go. **⭐36**
* [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. **⭐2107**
* [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. **⭐3885**
* [jsonql](https://github.com/elgs/jsonql) - JSON query expression library in Golang. **⭐178**
* [jsonslice](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters. **⭐9**
* [rql](https://github.com/a8m/rql) - Resource Query Language for REST API. **⭐75**

## Resource Embedding

* [esc](https://github.com/mjibson/esc) - Embeds files into Go programs and provides http.FileSystem interfaces to them. **⭐364**
* [fileb0x](https://github.com/UnnoTed/fileb0x) - Simple tool to embed files in go with focus on "customization" and ease to use. **⭐352**
* [go-embed](https://github.com/pyros2097/go-embed) - Generates go code to embed resource files into your library or executable. **⭐12**
* [go-resources](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go. **⭐146**
* [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.
* [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries. **⭐1372**
* [statics](https://github.com/go-playground/statics) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. **⭐49**
* [statik](https://github.com/rakyll/statik) - Embeds static files into a Go executable. **⭐1566**
* [templify](https://github.com/wlbr/templify) - Embed external template files into Go code to create single file binaries. **⭐15**
* [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem. **⭐416**

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types. **⭐519**
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for Go for machine-learning and statistics (similar to pandas). **⭐20**
* [evaler](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator. **⭐33**
* [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages. **⭐241**
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang. **⭐36**
* [go-dsp](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go. **⭐556**
* [go-fn](https://github.com/ematvey/go-fn) - Mathematical functions written in Go language, that are not covered by math pkg. **⭐10**
* [go-gt](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language. **⭐4**
* [gocomplex](https://github.com/varver/gocomplex) - Complex number library for the Go programming language. **⭐4**
* [goent](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures **⭐10**
* [gohistogram](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams. **⭐112**
* [gonum/mat64](https://github.com/gonum/matrix) - The general purpose package for matrix computation. Package mat64 provides basic linear algebra operations for float64 matrices. **⭐460**
* [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. **⭐974**
* [goraph](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorith visualization). **⭐535**
* [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. **⭐1167**
* [GoStats](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. **⭐6**
* [graph](https://github.com/yourbasic/graph) - Library of basic graph algorithms. **⭐172**
* [ode](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. **⭐7**
* [orb](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. **⭐62**
* [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go. **⭐41**
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library. **⭐5**
* [PiHex](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. **⭐4**
* [sparse](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. **⭐46**
* [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. **⭐945**
* [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. **⭐1310**
* [TextRank](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. **⭐53**
* [triangolatte](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. **⭐7**

## Security

*Libraries that are used to help make your application more secure.*

* [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. **⭐1622**
* [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. **⭐321**
* [argon2pw](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison. **⭐63**
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.
* [BadActor](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban. **⭐225**
* [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. **⭐1458**
* [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". **⭐105**
* [goArgonPass](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. **⭐7**
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords. **⭐22**
* [lego](https://github.com/xenolf/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). **⭐2975**
* [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. **⭐827**
* [nacl](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's. **⭐371**
* [passlib](https://github.com/hlandau/passlib) - Futureproof password hashing library. **⭐205**
* [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. **⭐1049**
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in. **⭐142**
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys. **⭐171**

## Serialization

*Libraries and tools for binary serialization.*

* [asn1](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang. **⭐33**
* [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go. **⭐59**
* [colfer](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format. **⭐409**
* [csvutil](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures. **⭐246**
* [fwencoder](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go. **⭐6**
* [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. **⭐266**
* [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. **⭐1040**
* [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. **⭐2140**
* [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. **⭐3689**
* [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". **⭐3867**
* [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. **⭐1802**
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. **⭐102**
* [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures. **⭐77**
* [structs](https://github.com/fatih/structs) - Library with support for converting structs to maps, struct keys/values to slices, and more. **⭐2499**

## Server Applications

* [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. **⭐486**
* [Caddy](https://github.com/mholt/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. **⭐19184**
* [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
* [devd](https://github.com/cortesi/devd) - Local webserver for developers. **⭐2640**
* [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. **⭐305**
* [etcd](https://github.com/coreos/etcd) - Highly-available key value store for shared configuration and service discovery. **⭐21248**
* [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. **⭐534**
* [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. **⭐485**
* [jackal](https://github.com/ortuman/jackal) - An XMPP server written in Go. **⭐514**
* [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. **⭐13588**
* [nsq](http://nsq.io/) - A realtime distributed messaging platform.
* [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. **⭐1765**
* [yakvs](https://git.sci4me.com/sci4me/yakvs) - Small, networked, in-memory key-value store.

## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace) - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. **⭐731**
* [amber](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. **⭐798**
* [damsel](https://github.com/dskinner/damsel) - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. **⭐20**
* [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. **⭐391**
* [extemplate](https://github.com/dannyvankooten/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance. **⭐5**
* [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). **⭐228**
* [gofpdf](https://github.com/jung-kurt/gofpdf) - PDF document generator with high level support for text, drawing and images. **⭐2590**
* [hero](https://github.com/shiyanhui/hero) - Hero is a handy, fast and powerful go template engine. **⭐1069**
* [jet](https://github.com/CloudyKit/jet) - Jet template engine. **⭐515**
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation.
* [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates. **⭐68**
* [mustache](https://github.com/hoisie/mustache) - Go implementation of the Mustache template language. **⭐945**
* [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. **⭐1338**
* [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. **⭐1146**
* [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go. **⭐292**
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang. **⭐641**
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). **⭐134**
* [velvet](https://github.com/gobuffalo/velvet) - Complete handlebars implementation in Go. **⭐62**

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. **⭐10**
    * [badio](https://github.com/cavaliercoder/badio) - Extensions to Go's `testing/iotest` package. **⭐8**
    * [baloo](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy. **⭐553**
    * [biff](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible. **⭐5**
    * [bro](https://github.com/marioidival/bro) - Watch files in directory and run tests for them. **⭐22**
    * [charlatan](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests. **⭐179**
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework. **⭐58**
    * [dbcleaner](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby. **⭐49**
    * [dsunit](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files. **⭐20**
    * [endly](https://github.com/viant/endly) - Declarative end to end functional testing. **⭐54**
    * [frisby](https://github.com/verdverm/frisby) - REST API testing framework. **⭐233**
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal. **⭐181**
    * [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests **⭐700**
    * [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code. **⭐183**
    * [go-testdeep](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package. **⭐18**
    * [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. **⭐247**
    * [goblin](https://github.com/franela/goblin) - Mocha like testing framework fo Go. **⭐556**
    * [gocheck](http://labix.org/gocheck) - More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload. **⭐4000**
    * [gocrest](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions. **⭐7**
    * [godog](https://github.com/DATA-DOG/godog) - Cucumber or Behat like BDD framework for Go. **⭐478**
    * [gofight](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework. **⭐202**
    * [gogiven](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go. **⭐7**
    * [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
    * [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language. **⭐110**
    * [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. **⭐50**
    * [gosuite](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. **⭐7**
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. **⭐24**
    * [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. **⭐972**
    * [restit](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test. **⭐49**
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. **⭐219**
    * [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. **⭐6156**
    * [testsql](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished. **⭐4**
    * [Tt](https://github.com/vcaesar/tt) - Simple and colorful test tools. **⭐2**
    * [wstest](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler. **⭐48**

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. **⭐281**
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. **⭐1133**
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes. **⭐92**
    * [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. **⭐622**
    * [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language. **⭐1796**
    * [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing. **⭐49**
    * [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces. **⭐145**
    * [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter. **⭐22**

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. **⭐2424**
    * [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. **⭐456**
    * [Tavor](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework. **⭐189**

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. **⭐194**
    * [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. **⭐2432**
    * [ggr](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium Wedriver requests to multiple Selenium hubs. **⭐173**
    * [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. **⭐813**

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align) - A general purpose application that aligns text. **⭐45**
    * [allot](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots. **⭐31**
    * [bbConvert](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags. **⭐5**
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. **⭐3364**
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. **⭐1028**
    * [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers **⭐6265**
    * [commonregex](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go **⭐507**
    * [dataflowkit](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data. **⭐160**
    * [doi](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go. **⭐2**
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go. **⭐26**
    * [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). **⭐5**
    * [encdec](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decodersa. **⭐3**
    * [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings. **⭐48**
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection). **⭐14**
    * [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. **⭐1627**
    * [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language. **⭐66**
    * [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string. **⭐159**
    * [go-slugify](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support. **⭐23**
    * [go-vcard](https://github.com/emersion/go-vcard) - Parse and format vCard. **⭐13**
    * [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. **⭐932**
    * [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language. **⭐221**
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - Format bytes to string.
    * [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts. **⭐27**
    * [goq](https://github.com/andrewstuart/goq) - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). **⭐117**
    * [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. **⭐6480**
    * [goregen](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions. **⭐32**
    * [gotext](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go. **⭐181**
    * [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text. **⭐39**
    * [htmlquery](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression **⭐35**
    * [inject](https://github.com/facebookgo/inject) - Package inject provides a reflect based injector. **⭐1018**
    * [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. **⭐293**
    * [sdp](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. **⭐45**
    * [sh](https://github.com/mvdan/sh) - Shell parser and formatter. **⭐1438**
    * [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. **⭐240**
    * [Slugify](https://github.com/avelino/slugify) - Go slugify application that handles string. **⭐25**
    * [syndfeed](https://github.com/zhengchun/syndfeed) - A syndication feed for Atom 1.0 and RSS 2.0. **⭐4**
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). **⭐2319**
* Utility
    * [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go. **⭐186**
    * [kace](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. **⭐8**
    * [parseargs-go](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes. **⭐4**
    * [parth](https://github.com/codemodus/parth) - URL path segmentation parsing. **⭐23**
    * [radix](https://github.com/yourbasic/radix) - fast string sorting algorithm. **⭐56**
    * [xj2go](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct. **⭐14**
    * [xurls](https://github.com/mvdan/xurls) - Extract urls from text. **⭐378**

## Third-party APIs

*Libraries for accessing third party APIs.*

* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). **⭐32**
* [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. **⭐927**
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) - The official AWS SDK for the Go programming language. **⭐4307**
* [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API. **⭐14**
* [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/). **⭐57**
* [circleci](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API. **⭐32**
* [clarifai](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API. **⭐58**
* [codeship-go](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2. **⭐15**
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API **⭐3**
* [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. **⭐725**
* [ethrpc](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API. **⭐142**
* [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. **⭐690**
* [fcm](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging. **⭐27**
* [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API. **⭐36**
* [gami](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface. **⭐26**
* [gcm](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging. **⭐29**
* [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. **⭐260**
* [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. **⭐4079**
* [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. **⭐344**
* [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler **⭐3**
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API. **⭐7**
* [go-imgur](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com) **⭐9**
* [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) **⭐346**
* [go-marathon](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS. **⭐177**
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). **⭐14**
* [go-sophos](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. **⭐2**
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) - Go client library for the SPTrans Olho Vivo API. **⭐5**
* [go-telegraph](https://github.com/toby3d/go-telegraph) - Telegraph publishing platform API client. **⭐48**
* [go-tgbot](https://github.com/olebedev/go-tgbot) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. **⭐78**
* [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. **⭐95**
* [go-twitch](https://github.com/knspriggs/go-twitch) - Go client for interacting with the Twitch v3 API. **⭐14**
* [go-twitter](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs. **⭐545**
* [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API. **⭐13**
* [go-xkcd](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API. **⭐36**
* [goamz](https://github.com/mitchellh/goamz) - Popular fork of [goamz](https://launchpad.net/goamz) which adds some missing API calls to certain packages. **⭐680**
* [golyrics](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website. **⭐28**
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library. **⭐31**
* [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. **⭐1588**
* [google-analytics](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting. **⭐12**
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. **⭐1426**
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). **⭐5**
* [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. **⭐115**
* [govkbot](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library. **⭐18**
* [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. **⭐108**
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP.
* [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/). **⭐18**
* [Medium](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API. **⭐105**
* [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. **⭐54**
* [micha](https://github.com/onrik/micha) - Go Library for Telegram bot api. **⭐8**
* [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. **⭐555**
* [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. **⭐25**
* [patreon-go](https://github.com/mxpv/patreon-go) - Go library for Patreon API. **⭐14**
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API. **⭐259**
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK. **⭐1**
* [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API. **⭐48**
* [rrdaclient](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. **⭐8**
* [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API. **⭐19**
* [slack](https://github.com/nlopes/slack) - Slack API in Go. **⭐1901**
* [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API. **⭐9**
* [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API. **⭐16**
* [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers. **⭐11**
* [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. **⭐807**
* [tbot](https://github.com/yanzay/tbot) - Telegram bot server with API similar to net/http. **⭐167**
* [telebot](https://github.com/tucnak/telebot) - Telegram bot framework written in Go. **⭐747**
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client. **⭐1254**
* [textbelt](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API. **⭐12**
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) - Simple golang package to communicate with [themoviedb.org](https://themoviedb.org). **⭐11**
* [translate](https://github.com/poorny/translate) - Go online translation package. **⭐25**
* [Trello](https://github.com/adlio/trello) - Go wrapper for the Trello API. **⭐72**
* [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API. **⭐6**
* [uptimerobot](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API. **⭐24**
* [webhooks](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket. **⭐241**
* [wit-go](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API. **⭐14**
* [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API
* [zooz](https://github.com/gojuno/go-zooz) - Go client for the Zooz API. **⭐4**

## Utilities

*General utilities and tools to make your life easier.*

* [abutil](https://github.com/bahlo/abutil) - Collection of often-used Golang helpers. **⭐48**
* [apm](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API. **⭐117**
* [backscanner](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. **⭐6**
* [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. **⭐793**
* [chyle](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities. **⭐95**
* [circuit](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. **⭐50**
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. **⭐694**
* [clockwerk](http://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax.
* [clockwork](https://github.com/whiteShtef/clockwork) - Simple and intuitive job scheduling library in Go. **⭐33**
* [command](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher. **⭐9**
* [copy-pasta](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage. **⭐30**
* [ctop](https://github.com/bcicen/ctop) - [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. **⭐7596**
* [ctxutil](https://github.com/posener/ctxutil) - A collection of utility functions for contexts. **⭐0**
* [Death](https://github.com/vrecan/death) - Managing go application shutdown with signals. **⭐103**
* [Deepcopier](https://github.com/ulule/deepcopier) - Simple struct copying for Go. **⭐174**
* [delve](https://github.com/derekparker/delve) - Go debugger. **⭐9758**
* [dlog](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls. **⭐15**
* [ergo](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy. **⭐236**
* [evaluator](https://github.com/nullne/evaluator) - Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. **⭐10**
* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - Golang library for reading and writing Microsoft Excel™ (XLSX) files. **⭐2687**
* [fastlz](https://github.com/digitalcrab/fastlz) - Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang. **⭐10**
* [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. **⭐380**
* [filler](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag. **⭐14**
* [filter](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data. **⭐4**
* [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. **⭐17962**
* [gaper](https://github.com/maxcnunes/gaper) - Builds and restarts a Go project when it crashes or some watched file changes. **⭐29**
* [generate](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex. **⭐12**
* [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. **⭐567**
* [git-time-metric](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git. **⭐617**
* [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. **⭐3593**
* [go-astitodo](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code. **⭐41**
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only). **⭐158**
* [go-cron](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. **⭐168**
* [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go. **⭐413**
* [go-excel](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table. **⭐30**
* [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). **⭐592**
* [go-health](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services. **⭐57**
* [go-httpheader](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields. **⭐14**
* [go-rate](https://github.com/beefsack/go-rate) - Timed rate limiter for Go. **⭐271**
* [go-respond](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses. **⭐13**
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go. **⭐86**
* [go-torch](https://github.com/uber/go-torch) - Stochastic flame graph profiler for Go programs. **⭐3358**
* [go-trigger](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. **⭐158**
* [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package. **⭐38**
* [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons. **⭐375**
* [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. **⭐3610**
* [gohper](https://github.com/cosiner/gohper) - Various tools/modules help for development. **⭐246**
* [gojq](https://github.com/elgs/gojq) - JSON query in Golang. **⭐128**
* [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. **⭐1838**
* [golarm](https://github.com/msempere/golarm) - Fire alarms with system events. **⭐32**
* [golog](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks. **⭐41**
* [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs. **⭐414**
* [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images. **⭐20**
* [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. **⭐3466**
* [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. **⭐2236**
* [goreq](https://github.com/franela/goreq) - Minimal and simple request library for Go language. **⭐690**
* [goreq](https://github.com/smallnest/goreq) - Enhanced simplified HTTP client based on gorequest. **⭐83**
* [gorequest](https://github.com/parnurzeal/gorequest) - Simplified HTTP client with rich features for Go. **⭐1968**
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features. **⭐15**
* [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go. **⭐112**
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. **⭐12**
* [gpath](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection. **⭐25**
* [grequests](https://github.com/levigross/grequests) - Elegant and simple `net/http` wrapper that follows Python's requests library. **⭐1174**
* [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. **⭐531**
* [gubrak](https://gubrak.github.io/) - Golang utility library with syntactic sugar. It's like lodash, but for golang.
* [htcat](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility. **⭐461**
* [httpcontrol](https://github.com/facebookgo/httpcontrol) - Package httpcontrol allows for HTTP transport level control around timeouts and retries. **⭐475**
* [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. **⭐14017**
* [hystrix-go](https://github.com/afex/hystrix-go) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. **⭐1430**
* [immortal](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor. **⭐542**
* [intrinsic](https://github.com/mengzhuo/intrinsic) - Use x86 SIMD without writing any assembly code. **⭐35**
* [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. **⭐491**
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference. **⭐4**
* [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON. **⭐53**
* [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. **⭐83**
* [jsonhal](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses. **⭐8**
* [kazaam](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents. **⭐97**
* [leprechaun](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling **⭐6**
* [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go. **⭐97**
* [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. **⭐858**
* [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. **⭐621**
* [mimemagic](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility. **⭐30**
* [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. **⭐44**
* [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. **⭐1599**
* [minquery](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). **⭐42**
* [mmake](https://github.com/tj/mmake) - Modern Make. **⭐1431**
* [moldova](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template. **⭐146**
* [mp](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON. **⭐29**
* [mssqlx](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. **⭐47**
* [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers. **⭐57**
* [myhttp](https://github.com/inancgumus/myhttp) - Simple API to make HTTP GET requests with timeout support. **⭐27**
* [netbug](https://github.com/e-dard/netbug) - Easy remote profiling of your services. **⭐65**
* [okrun](https://github.com/xta/okrun) - go run error steamroller. **⭐13**
* [onecache](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). **⭐79**
* [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. **⭐1799**
* [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. **⭐4945**
* [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency. **⭐282**
* [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API. **⭐61**
* [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. **⭐832**
* [rclient](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs. **⭐21**
* [realize](https://github.com/tockins/realize) - Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. **⭐2596**
* [repeat](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating. **⭐51**
* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. **⭐328**
* [rerate](https://github.com/abo/rerate) - Redis-based rate counter and rate limiter for Go. **⭐11**
* [rerun](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes. **⭐150**
* [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. **⭐1364**
* [retry](https://github.com/kamilsk/retry) - Functional mechanism based on context to perform actions repetitively until successful. **⭐68**
* [retry](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go. **⭐2**
* [retry](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go. **⭐27**
* [retry](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done. **⭐7**
* [retry-go](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang. **⭐24**
* [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics. **⭐129**
* [rq](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client. **⭐23**
* [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy. **⭐249**
* [sling](https://github.com/dghubble/sling) - Go HTTP requests builder for API clients. **⭐752**
* [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. **⭐609**
* [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. **⭐5247**
* [sslice](https://github.com/yaa110/sslice) - Create a slice which is always sorted. **⭐1**
* [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. **⭐1152**
* [structs](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs. **⭐8**
* [Task](https://github.com/go-task/task) - simple "Make" alternative. **⭐1045**
* [toolbox](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. **⭐41**
* [ugo](https://github.com/alxrm/ugo) - ugo is slice toolbox with concise syntax for Go. **⭐19**
* [UNIS](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go. **⭐67**
* [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. **⭐4395**
* [util](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...). **⭐90**
* [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. **⭐7735**
* [xferspdy](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang. **⭐63**
* [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. **⭐2799**
* [xlsx](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs. **⭐26**

## Validation

*Libraries for validation.*

* [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. **⭐2955**
* [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. **⭐510**
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. **⭐808**
* [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. **⭐16**
* [validate](https://github.com/markbates/validate) - This package provides a framework for writing validations for Go applications. **⭐38**
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. **⭐2369**

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. **⭐62**
* [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. **⭐1226**
* [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go. **⭐65**
* [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. **⭐12**

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries. **⭐416**
* [go-astisub](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). **⭐131**
* [go-astits](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO. **⭐220**
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) - Parser and generator library for Apple m3u8 playlists **⭐17**
* [goav](https://github.com/giorgisio/goav) - Comphrensive Go bindings for FFmpeg. **⭐546**
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer. **⭐146**
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) - Subtitle format support for go. Supports .srt, .ttml, and .ass. **⭐9**
* [libvlc-go](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). **⭐32**
* [v4l](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go. **⭐20**

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
* [Aero](https://github.com/aerogo/aero) - High-performance web framework for Go, reaches top scores in Lighthouse. **⭐73**
* [Air](https://github.com/aofei/air) - An ideally refined web framework for Go. **⭐94**
* [Banjo](https://github.com/nsheremet/banjo) - Very simple and fast web framework for Go. **⭐3**
* [Beego](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language. **⭐17772**
* [Buffalo](http://gobuffalo.io) - Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. **⭐12151**
* [Fireball](https://github.com/zpatrick/fireball) - More "natural" feeling web framework. **⭐41**
* [Florest](https://github.com/jabong/florest-core) - High-performance workflow based REST API framework.
* [Gem](https://github.com/go-gem/gem) - Simple and fast web framework, friendly to REST API. **⭐153**
* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. **⭐22012**
* [Gizmo](https://github.com/NYTimes/gizmo) - Microservice toolkit used by the New York Times. **⭐2445**
* [go-json-rest](https://github.com/ant0ine/go-json-rest) - Quick and easy way to setup a RESTful JSON API. **⭐3200**
* [go-relax](https://github.com/codehack/go-relax) - Framework of pluggable components to build RESTful API's. **⭐153**
* [go-rest](https://github.com/ungerik/go-rest) - Small and evil REST framework for Go. **⭐110**
* [goa](https://github.com/raphael/goa) - Framework for developing microservices based on the design of Ruby's Praxis. **⭐3123**
* [Golax](https://github.com/fulldump/golax) - A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. **⭐64**
* [Golf](https://github.com/dinever/golf) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. **⭐229**
* [Gondola](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster. **⭐312**
* [gongular](https://github.com/mustafaakin/gongular) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection. **⭐401**
* [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support. **⭐45**
* [Macaron](https://github.com/go-macaron/macaron) - Macaron is a high productive and modular design web framework in Go. **⭐2581**
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. **⭐330**
* [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang. **⭐45**
* [neo](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API. **⭐378**
* [Resoursea](https://github.com/resoursea/api) - REST framework for quickly writing resource based services. **⭐29**
* [REST Layer](http://rest-layer.io) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. **⭐10465**
* [rex](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. **⭐24**
* [sawsij](https://github.com/jaybill/sawsij) - lightweight, open-source web framework for building high-performance, data-driven web applications. **⭐2**
* [tango](https://github.com/lunny/tango) - Micro & pluggable web framework for Go. **⭐743**
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - Go framework for building JSON web services inspired by Dropwizard. **⭐989**
* [traffic](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go. **⭐514**
* [utron](https://github.com/gernest/utron) - Lightweight MVC framework for Go(Golang). **⭐2094**
* [vox](https://github.com/aisk/vox) - A golang web framework for humans, inspired by Koa heavily. **⭐12**
* [WebGo](https://github.com/bnkamalesh/webgo) - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). **⭐55**
* [YARF](https://github.com/yarf-framework/yarf) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way. **⭐45**
* [Zerver](https://github.com/cosiner/zerver) - Zerver is an expressive, modular, feature completed RESTful framework. **⭐144**

### Middlewares

#### Actual middlewares

* [client-timing](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header. **⭐10**
* [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. **⭐929**
* [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST. **⭐30**
* [go-server-timing](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header. **⭐716**
* [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. **⭐416**
* [ln-paywall](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin) **⭐66**
* [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. **⭐993**
* [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends. **⭐68**

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. **⭐1674**
* [catena](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain"). **⭐7**
* [chain](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware"). **⭐63**
* [go-wrap](https://github.com/go-on/wrap) - Small middlewares package for net/http. **⭐54**
* [gores](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. **⭐77**
* [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang. **⭐287**
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http. **⭐203**
* [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. **⭐5896**
* [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. **⭐1165**
* [renderer](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. **⭐112**
* [rye](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. **⭐88**
* [stats](https://github.com/thoas/stats) - Go middleware that stores various information about your web application. **⭐504**
* [Volatile](https://github.com/volatile/core) - Minimalist middleware stack promoting flexibility, good practices and clean code. **⭐81**

### Routers

* [alien](https://github.com/gernest/alien) - Lightweight and fast http router from outer space. **⭐95**
* [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. **⭐1171**
* [Bxog](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. **⭐90**
* [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. **⭐4486**
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`. **⭐590**
* [FastRouter](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go. **⭐17**
* [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. **⭐1348**
* [Goji](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. **⭐679**
* [GoRouter](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. **⭐36**
* [gowww/router](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface. **⭐150**
* [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. **⭐8081**
* [httptreemux](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. **⭐358**
* [lars](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. **⭐361**
* [medeina](https://github.com/imdario/medeina) - Medeina is a HTTP routing tree based on HttpRouter, inspired by Roda and Cuba. **⭐18**
* [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. **⭐7504**
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. **⭐308**
* [pat](https://github.com/bmizerany/pat) - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra. **⭐1150**
* [pure](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation. **⭐67**
* [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers. **⭐348**
* [vestigo](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications. **⭐234**
* [violetear](https://github.com/nbari/violetear) - Go HTTP router. **⭐91**
* [xmux](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support. **⭐86**
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go. **⭐343**
* [zeus](https://github.com/daryl/zeus) - Very simple and fast HTTP router for Go. **⭐111**

## Windows

* [d3d9](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9. **⭐74**
* [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. **⭐453**

## XML

*Libraries and tools for manipulating XML.*

* [XML-Comp](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags. **⭐15**
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module. **⭐5**
* [xpath](https://github.com/antchfx/xpath) - XPath package for Go. **⭐94**
* [xquery](https://github.com/antchfx/xquery) - XQuery lets you extract data from HTML/XML documents using XPath expression. **⭐137**
* [xml2map](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang **⭐7**
# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes. **⭐153**
* [dupl](https://github.com/mibk/dupl) - Tool for code clone detection. **⭐128**
* [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. **⭐1139**
* [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time. **⭐848**
* [Go Metalinter](https://github.com/alecthomas/gometalinter) - Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form.
* [go-checkstyle](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments. **⭐85**
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. **⭐224**
* [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. **⭐405**
* [go-outdated](https://github.com/firstrow/go-outdated) - Console application that displays outdated packages. **⭐46**
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer. **⭐307**
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/) - GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code. **⭐2749**
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - gosimple is a linter for Go source code that specialises on simplifying code.
* [gostatus](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages. **⭐226**
* [interfacer](https://github.com/mvdan/interfacer) - Linter that suggests interface types. **⭐719**
* [lint](https://github.com/surullabs/lint) - Run linters as part of go test. **⭐61**
* [php-parser](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go. **⭐530**
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) - tarp finds functions and methods without direct unit tests in Go source code. **⭐11**
* [unconvert](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source. **⭐227**
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - unused checks Go code for unused constants, variables, functions and types.
* [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags. **⭐62**

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. **⭐17**
* [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs.
* [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. **⭐1422**
* [Goclipse](https://github.com/GoClipse/goclipse) - Eclipse plugin for Go. **⭐776**
* [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. **⭐4499**
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. **⭐3054**
* [gounit-vim](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature. **⭐14**
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) - Go language support for the Theia IDE. **⭐8**
* [velour](https://github.com/velour/velour) - IRC client for the acme editor. **⭐16**
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save. **⭐77**
* [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. **⭐9415**
* [vscode-go](https://github.com/Microsoft/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. **⭐4185**
* [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes. **⭐153**

## Go Generate Tools

* [generic](https://github.com/usk81/generic) - flexible data type for Go. **⭐22**
* [genny](https://github.com/cheekybits/genny) - Elegant generics for Go. **⭐745**
* [gocontracts](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation. **⭐27**
* [gonerics](http://github.com/bouk/gonerics) - Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. **⭐1503**
* [gounit](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates. **⭐14**
* [re2dfa](https://github.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code. **⭐160**

## Go Tools

* [colorgo](https://github.com/songgao/colorgo) - Wrapper around `go` command for colorized `go build` output. **⭐94**
* [depth](https://github.com/KyleBanks/depth) - Visualize dependency trees of any package by analyzing imports. **⭐270**
* [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) - A [Yeoman](http://yeoman.io) generator to get new Go projects started. **⭐8**
* [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. **⭐1337**
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) - Bash completion for go and wgo. **⭐37**
* [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. **⭐2904**
* [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. **⭐3040**
* [richgo](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations. **⭐306**
* [rts](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses. **⭐177**

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool. **⭐1638**
* [aurora](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console. **⭐328**
* [awsenv](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile. **⭐16**
* [Banshee](https://github.com/eleme/banshee) - Anomalies detection system for periodic metrics. **⭐998**
* [Blast](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs. **⭐152**
* [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. **⭐1182**
* [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. **⭐2677**
* [DepCharge](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects. **⭐6**
* [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). **⭐196**
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI. **⭐16**
* [drone-scp](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI. **⭐41**
* [Dropship](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn. **⭐42**
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. **⭐67**
* [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts **⭐1508**
* [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. **⭐2474**
* [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. **⭐10836**
* [Go Metrics](https://github.com/rcrowley/go-metrics) - Go port of Coda Hale's Metrics library: https://github.com/codahale/metrics.
* [go-furnace](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. **⭐50**
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. **⭐623**
* [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go. **⭐174**
* [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application. **⭐216**
* [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. **⭐298**
* [govvv](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries. **⭐286**
* [gox](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool. **⭐2968**
* [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging. **⭐1575**
* [grapes](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease. **⭐113**
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. **⭐3910**
* [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. **⭐4313**
* [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. **⭐1273**
* [kcli](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages. **⭐47**
* [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. **⭐43942**
* [lstags](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries. **⭐192**
* [lwc](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command. **⭐5**
* [manssh](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily. **⭐184**
* [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. **⭐51122**
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data. **⭐253**
* [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. **⭐157**
* [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. **⭐8188**
* [Pewpew](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester. **⭐171**
* [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies. **⭐31**
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. **⭐951**
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker). **⭐467**
* [sg](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. **⭐3**
* [skm](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! **⭐462**
* [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. **⭐1005**
* [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. **⭐18680**
* [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! **⭐9586**
* [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. **⭐3332**
* [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
* [winrm-cli](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines. **⭐45**

### Other Software
* [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. **⭐1380**
* [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine. **⭐70**
* [Cherry](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go. **⭐176**
* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. **⭐1694**
* [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. **⭐5842**
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. **⭐5612**
* [DDNS](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend. **⭐59**
* [Docker](http://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. **⭐524**
* [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. **⭐2386**
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH.
* [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. **⭐371**
* [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go. **⭐208**
* [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. **⭐289**
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list. **⭐12**
* [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. **⭐9627**
* [hugo](http://gohugo.io/) - Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go. **⭐238**
* [ipe](https://github.com/dimiro1/ipe) - Open source Pusher server implementation compatible with Pusher client libraries written in GO. **⭐240**
* [JayDiff](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go. **⭐25**
* [joincap](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together. **⭐107**
* [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Leaps](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms. **⭐615**
* [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. **⭐1489**
* [limetext](http://limetext.org/) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. **⭐4881**
* [mockingjay](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. **⭐373**
* [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. **⭐2091**
* [naclpipe](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go. **⭐20**
* [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. **⭐3825**
* [orange-cat](https://github.com/noraesae/orange-cat) - Markdown previewer written in Go. **⭐172**
* [Orbit](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates. **⭐114**
* [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. **⭐529**
* [Pipe](https://github.com/b3log/pipe) - A small and beautiful blogging platform. **⭐1600**
* [restic](https://github.com/restic/restic) - De-duplicating backup program. **⭐5497**
* [rkt](https://github.com/coreos/rkt) - App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM. **⭐8293**
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
* [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). **⭐289**
* [snap](https://github.com/intelsdi-x/snap) - Powerful telemetry framework. **⭐1783**
* [Snitch](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru. **⭐15**
* [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
* [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
* [Tenyks](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging. **⭐166**
* [term-quiz](https://github.com/crazcalm/term-quiz) - Quizzes for your terminal. **⭐14**
* [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. **⭐3207**
* [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. **⭐468**
* [wellington](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass). **⭐283**

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions. **⭐88**
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. **⭐18**
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. **⭐108**
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. **⭐1152**
* [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) - Naive performance test of two ways to do type assertion in Go. **⭐5**
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. **⭐825**
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. **⭐724**
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. **⭐51**
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) - Tiny collection of Go micro benchmarks. The intent is to compare some language features to others. **⭐15**
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities. **⭐47**
* [gospeed](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs. **⭐84**
* [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark. **⭐14**
* [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. **⭐872**
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language. **⭐154**

## Conferences

* [Capital Go](http://www.capitalgolang.com) - Washington, D.C., USA
* [dotGo](http://www.dotgo.eu) - Paris, France
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan
* [GoDays](https://www.godays.io/) - Berlin, Germany
* [GoLab](http://golab.io/) - Florence, Italy
* [GolangUK](http://golanguk.com/) - London, UK
* [GopherChina](http://gopherchina.org) - Shanghai, China
* [GopherCon](http://www.gophercon.com/) - Denver, USA
* [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, BR
* [GopherCon Europe](https://gophercon.is/) - Reykjavik, Iceland
* [GopherCon India](https://www.gophercon.in/) - Pune, India
* [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel
* [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia
* [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore
* [GothamGo](http://gothamgo.com/) - New York City, USA
* [GoWayFest](https://goway.io/) - Minsk, Belarus

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. **⭐5589**
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg] **⭐25**
* [gopher-logos](https://github.com/GolangUA/gopher-logos) - adorable gopher logos **⭐48**
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers) **⭐392**
* [gopher-vector](https://github.com/golang-samples/gopher-vector) **⭐318**
* [gophericons](https://github.com/shalakhin/gophericons) **⭐548**
* [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself
* [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara **⭐1505**
* [gophers](https://github.com/egonelbre/gophers) - Free gophers **⭐1294**
* [gophers](https://github.com/rogeralsing/gophers) - random gopher graphics **⭐47**
* [gophers](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern **⭐27**

## Meetups

* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
* [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
* [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
* [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*Add the group of your city/country here (send **PR**)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## Websites

* [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. **⭐22686**
* [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org) - The official Go blog.
* [Go Challenge](http://golang-challenge.org/) - Learn Go by solving problems and getting feedback from Go experts.
* [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com) - A report card for your Go package.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go. **⭐31**
* [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
* [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com) - Links and news about Go programming.
* [golang-graphics](https://github.com/mholt/golang-graphics) - Collection of Go images, graphics, and art. **⭐137**
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [r/Golang](https://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries
* [Games With Go](http://gameswithgo.org/) - A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card.
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8) - Interactively edit & play Go snippets on your mobile device
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
* [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development.
* [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
* [Working with Go](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers.
* [Your basic Go](http://yourbasic.org/golang) - Huge collection of tutorials and how to's
