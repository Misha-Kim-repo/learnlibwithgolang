# learnlibwithgolang

#도서: GO 프로그래밍 쿡북 2/e(https://www.aladin.co.kr/shop/wproduct.aspx?ISBN=K172836703&start=pnaver_02)

Chapter 01: I/O와 파일 시스템
 - I/O 인터페이스(io.Reader/io.Writer) 사용
 - bytes 및 strings 패키지 활용
 - encoding/csv을 활용한 CSV 작업(Read/Write)
 - ioutil 패키지를 활용하여 임시 파일 및 디렉터리 생성
 - text/template 및 html/template 패키지를 활용한 여러 템플릿 작성
Chapter 02: 명령줄 도구
 - flag 패키지를 활용한 커스텀 변수 타입/내장 변수/축약어 플래그/플래그 기록법 사용 등
 - 서드파티 패키지를 활용한 환경 변수 값 읽기/특정 요구 사항 지정하기
 - TOML/YAML/JSON 형식의 파일 형식 변환
 - os/signal 패키지를 통해 신호(signal) 전달
 - ANSI 컬러 설정 
Chapter 03: 데이터 변환 및 구성
 - strconv 패키지와 인터페이스 리플렉션을 사용하여 서로 다른 타입 변환을 처리
 - math 패키지를 사용하여 복잡한 부동 소수점 연산 처리
 - math/big 패키지를 통해 64비트 이상의 숫자들을 처리
 - sql.NullType64 타입 사용 방식
 - gob 패키지를 활용, Go 데이터 타입을 처리
 - base64 패키지를 활용하여 GET 요청 등의 데이터 처리 및 문자열 표현 인코딩 생성
 - reflect 패키지를 사용하여 리플렉션 처리
 - Go 클로저를 통해 컬렉션 생성
Chapter 04: Go의 오류 처리
  - Error 인터페이스를 통해 기본적인 오류 처리
  - pkg/errors 패키지를 통해 디폴트 errors 패키지를 대체하고 반환되는 오류를 래핑하여 추가 정보 제공
  - log 패키지를 사용하여 예상치 못한 동작에 대한 로깅 수행
  - apex/log 및 sirupsen/logrus 패키지를 활용하여 구조화된 로깅 수행 및 서드 파티에 맞는 포맷으로 변경
  - pkg/context 패키지를 활용하여 함수 사이에 변수를 추가로 전달하여 풍부한 정보를 로깅에 제공
  - panic 문제에 대해 파악
Chapter 05: 네트워크 프로그래밍
  - net 패키지를 통해 TCP/UDP 연결 처리 및 DNS 기능 구현
  - github.com/gorilla/websocket 패키지를 통해 웹소켓 작업 구현
  - net/rpc 패키지를 활용하여 원격 메소드 호출
  - net/mail 패키지를 활용하여 이메일 원시 텍스트를 파싱
Chapter 06: 데이터베이스와 저장소의 모든 것
  - MYSQL과 함께 database/sql 패키지 활용하여 연결 풀/연결 시간 자동화 및 트랜잭션 처리
  - redis를 활용한 작업 처리
  - MongoDB를 활용하여 NoSQL 사용/관련하여 저장소 인터페이스 생성
Chpater 07: 웹 클라이언트와 API
  - net/http 패키지를 활용하여 HTTP API 작업 수행 및 REST API 애플리케이션 작성
  - go routine을 활용하여 비동기 방식 애플리케이션을 구축하고 이를 REST API로 처리
  - oauth2 패키지를 화용하여 엔드포인트 지정 및 github과의 연동 작업 
  - (진행중)
