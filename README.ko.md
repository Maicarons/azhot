<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="로고" width="128" height="128" />
</p>

<p align="center">
  <img src="banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

[![Go 버전](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![라이선스](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![빌드 상태](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Go 리포트 카드](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

> 주요 플랫폼의 인기 검색 API를 제공하는 통합 서비스

## 📖 목차

- [프로젝트 소개](#프로젝트-소개)
- [기능 특징](#기능-특징)
- [지원 플랫폼](#지원-플랫폼)
- [빠른 시작](#빠른-시작)
- [API 사용법](#api-사용법)
- [MCP 서버](#mcp-서버)
- [개발 및 기여](#개발-및-기여)
- [라이선스](#라이선스)
- [문제 피드백](#문제-피드백)

## 프로젝트 소개

`azhot`은 주요 플랫폼의 인기 검색 데이터를 통합하는 API 서비스로, 다양한 플랫폼의 인기 검색 콘텐츠에 통합 인터페이스를 통해 접근할 수 있도록 제공합니다. 이 프로젝트는 Go 언어로 개발되었으며, Fiber 프레임워크를 기반으로 구축되어 주요 플랫폼의 실시간 인기 검색 순위 데이터를 가져오는 것을 지원합니다.

## 기능 특징

- 🚀 주요 플랫폼의 인기 검색 데이터를 가져오는 통합 API 인터페이스
- ⚡ 고성능, `Go`+`Fiber v2`로 개발되어 네이티브 캐시 메커니즘 + 접근 제어 기능 포함
- 🔄 데이터베이스에 인기 검색 데이터를 주기적으로 업데이트 [SQLite + MySQL + 확장 가능한 기타 DB 지원]
- 📚 [Swagger API 문서](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 RESTful API 설계
- 📦 예제 [프론트엔드](/frontend) 포함
- 🔌 WebSocket 실시간 데이터 푸시 지원
- 🤖 **신규** AI 모델 컨텍스트 프로토콜(MCP) 서버 지원

## 프로젝트 구조
```
azhot/
├── all/                 # all 기능 코드
├── app/                 # 주요 프로그램 코드
├── config/              # 설정 파일 읽기
├── docs/                # swagger API 문서
├── model/               # 데이터베이스 모델
├── mcp/                 # AI 모델 컨텍스트 프로토콜 서버
├── router/              # 라우팅 설정
├── service/             # 비즈니스 로직
├── websocket/           # WebSocket 기능
├── frontend/            # 템플릿 파일
├── .env                 # 환경 변수
├── Dockerfile           # Docker 빌드 파일
├── go.mod               # Go 모듈 정의
├── main.go              # 주요 프로그램 파일
└── README.md            # 프로젝트 설명 문서
```

## 지원 플랫폼

| 이름 | 라우트 이름 | 사용 가능 여부 |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| 360 검색 | 360search | ✅ |
| AcFun | acfun | ✅ |
| 바이두 | baidu | ✅ |
| 빌리빌리 | bilibili | ✅ |
| CCTV | cctv | ✅ |
| CSDN | csdn | ✅ |
| 동치우디 | dongqiudi | ✅ |
| 더우판 | douban | ✅ |
| 더우인 | douyin | ✅ |
| 깃허브 | github | ✅ |
| 내셔널 지오그래픽 | guojiadili | ✅ |
| 오늘의 역사 | historytoday | ✅ |
| 후푸 | hupu | ✅ |
| IT 홈 | ithome | ✅ |
| 리스핀 | lishipin | ✅ |
| 남방위크리 | nanfang | ✅ |
| 펑페이 뉴스 | pengpai | ✅ |
| 텐센트 뉴스 | qqnews | ✅ |
| 콰크 | quark | ✅ |
| 인민일보 온라인 | renmin | ✅ |
| 소구 | sougou | ✅ |
| 소후 | souhu | ✅ |
| 토토판 | toutiao | ✅ |
| V2EX | v2ex | ✅ |
| 네이트 뉴스 | wangyinews | ✅ |
| 웨이보 | weibo | ✅ |
| 신징바오 | xinjingbao | ✅ |
| 지후 | zhihu | ✅ |

## 빠른 시작

### 환경 요구 사항

- Go >= 1.18
- MySQL (선택 사항, 데이터 저장용)

### 설치 단계

1. 프로젝트 복제
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. 종속성 설치
```bash
go mod tidy
```

3. 환경 설정
```bash
# 설정 파일 복사
cp .env.example .env
# 설정 파일 편집
vim .env
```

4. API 문서 생성
```bash
swag init
```

5. 프로젝트 실행
```bash
# 개발 모드로 실행
make dev

# 또는 빌드 후 실행
make run
```

### Docker를 사용한 실행

```bash
# 이미지 빌드
docker build -t azhot .

# 컨테이너 실행
docker run -d -p 8080:8080 azhot
```

## API 사용법

### HTTP API

#### 모든 플랫폼 목록 가져오기

```
GET /list
```

모든 지원되는 플랫폼 정보를 가져옵니다.

#### 특정 플랫폼 인기 검색 가져오기

```
GET /{platform}
```

예를 들어, 지후 인기 검색을 가져오려면:
```
GET /zhihu
```

### WebSocket API

이 프로젝트는 WebSocket 실시간 데이터 푸시를 지원하며, HTTP API와 동일한 라우팅 구조를 제공합니다.

#### 일반 WebSocket 엔드포인트

```
ws://localhost:8080/ws
```

연결 후 특정 플랫폼 데이터를 구독하거나 요청하는 메시지를 보낼 수 있습니다.

#### 특정 플랫폼 WebSocket 엔드포인트

```
ws://localhost:8080/ws/{platform}
```

예를 들어, 바이두 인기 검색 WebSocket에 연결하려면:
```
ws://localhost:8080/ws/baidu
```

#### WebSocket 메시지 형식

```json
{
  "type": "subscribe|request|ping",
  "source": "플랫폼 이름, 예: baidu, zhihu 등",
  "data": {}
}
```

- `subscribe`: 특정 플랫폼의 실시간 데이터 구독
- `request`: 일회성 데이터 요청
- `ping`: 하트비트 메시지

#### WebSocket 엔드포인트 목록

- 일반 엔드포인트: `ws://localhost:8080/ws`
- 바이두: `ws://localhost:8080/ws/{platform}`
- 모든 플랫폼 통합: `ws://localhost:8080/ws/all`
- 플랫폼 목록: `ws://localhost:8080/ws/list`
- 역사적 조회 API:
  - `ws://localhost:8080/ws/history/{source}` - 지정된 플랫폼의 모든 역사적 데이터 가져오기
  - `ws://localhost:8080/ws/history/{source}/{date}` - 지정된 플랫폼, 날짜의 모든 시간 데이터 가져오기
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - 지정된 플랫폼, 날짜 및 시간의 역사적 데이터 가져오기
- 그리고 HTTP API에 해당하는 모든 WebSocket 엔드포인트

### API 응답 형식

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "2026년 신년 인사말",
      "url": "https://www.zhihu.com/search?q=2026년 신년 인사말"
    },
    // ...
    {
      "index": 12,
      "title": "동북 지역 네티즌, '샤오비가' 쥐 발견",
      "url": "https://www.zhihu.com/search?q=동북 지역 네티즌, '샤오비가' 쥐 발견"
    }
  ]
}
```

## MCP 서버

이 프로젝트는 이제 표준화된 프로토콜을 통해 AI 모델 및 스마트 어시스턴트가 인기 검색 데이터에 접근할 수 있도록 하는 AI 모델 컨텍스트 프로토콜(MCP) 서버를 통합했습니다.

### 기능 특징

- **표준화된 도구 인터페이스**: 표준 MCP 도구 목록 및 실행 인터페이스 제공
- **인기 검색 데이터 접근**: 도구를 통해 각 플랫폼의 인기 검색 데이터 가져오기 지원
- **역사적 데이터 조회**: 역사적 인기 검색 데이터 조회 지원
- **다양한 배포 모드**: HTTP 및 STDIO 두 가지 배포 모드 지원

### MCP 서버 활성화

`.env` 파일에서 다음 옵션을 구성하십시오:

```env
MCP_STDIO_ENABLED=true      # STDIO MCP 서버 활성화
MCP_HTTP_ENABLED=true       # HTTP MCP 서버 활성화
MCP_PORT=8081               # HTTP MCP 서버 포트
```

### MCP 도구 목록

- `get_hot_search`: 지정된 플랫폼의 인기 검색 데이터 가져오기
- `get_all_hot_search`: 모든 플랫폼의 인기 검색 데이터 통합 가져오기
- `get_history_data`: 지정된 플랫폼의 역사적 인기 검색 데이터 가져오기

### MCP 엔드포인트

- `/mcp/tools` - 사용 가능한 도구 목록 가져오기
- `/mcp/tool/execute` - 지정된 도구 실행
- `/mcp/prompts` - 사용 가능한 프롬프트 목록 가져오기
- `/mcp/ping` - 헬스 체크 엔드포인트
- `/mcp/.well-known/mcp-info` - MCP 서버 메타데이터

### 사용 예시

HTTP를 통해 MCP 도구 호출:
```bash
curl -X POST http://localhost:8080/mcp/tool/execute \
  -H "Content-Type: application/json" \
  -d '{
    "method": "tool/execute",
    "params": {
      "name": "get_hot_search",
      "arguments": {
        "platform": "zhihu"
      }
    },
    "id": "req-1",
    "jsonrpc": "2.0"
  }'
```

자세한 내용은 [MCP 서버 문서](mcp/README.md)를 참조하십시오.

## 개발 및 기여

모든 형태의 기여를 환영합니다! 프로젝트에 기여하고 싶다면 다음 단계를 따르십시오:

1. 이 프로젝트 포크
2. 기능 브랜치 생성 (`git checkout -b feature/AmazingFeature`)
3. 변경 사항 커밋 (`git commit -m 'Add some AmazingFeature'`)
4. 브랜치에 푸시 (`git push origin feature/AmazingFeature`)
5. 풀 리퀘스트 생성

### 로컬 개발

```bash
# 테스트 실행
dev.sh # Air를 핫 리로드 디버깅 도구로 사용
```

## CMake 빌드 시스템

이 프로젝트는 이제 CMake를 사용한 빌드를 지원하며, Windows 및 Linux 플랫폼을 지원합니다.

### 빌드 명령

```bash
# 현재 플랫폼용 빌드
mkdir build && cd build
cmake ..
cmake --build . --target build

# 실행
cmake --build . --target run

# 개발 모드로 실행
cmake --build . --target dev

# 크로스 플랫폼 빌드 (사전 정의된 플랫폼)
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# 크로스 플랫폼 빌드 (스크립트 사용)
# Linux/macOS:
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows:
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# 패키징 (지원되는 모든 플랫폼에 대해 zip 패키지 생성)
cmake --build . --target package

# 빌드 산물 정리
cmake --build . --target azhot_clean

# 테스트 실행
cmake --build . --target test

# 모든 테스트 실행
cmake --build . --target test-all

# 코드 포맷
cmake --build . --target fmt

# 종속성 정리
cmake --build . --target tidy

# 정적 분석
cmake --build . --target staticcheck

# CI 버전 빌드 (swagger 문서 생성 없음)
cmake --build . --target build-ci
```

## 라이선스

이 프로젝트는 AGPL-3.0 라이선스 하에 있습니다 - 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하십시오.

## 문제 피드백

사용 중 문제가 발생하거나 제안 사항이 있는 경우 이슈 또는 풀 리퀘스트를 제출해 주십시오.

- 🐛 [문제 보고](https://github.com/maicarons/azhot/issues)
- ✨ [기능 요청](https://github.com/maicarons/azhot/issues)

---

> 🌟 이 프로젝트가 도움이 되었다면 스타를 주세요! 이것이 우리에게 가장 큰 지원입니다!