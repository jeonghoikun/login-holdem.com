# webserver
매번 같은 웹서버 코드를 작성하기 때문에 초기화 소스를 미리 작성해둠<br>
모든 광고용 웹서버는 효율적인 유지보수를 위해 해당 소스를 기초로 작성함

1. $ git clone https://jeonghoikun:{{TOKEN}}@github.com/jeonghoikun/webserver.git
2. Directory 이름 변경 (ex. domain.com)
3. go.mod 파일에서 "webserver" 텍스트를 Directory 명으로 변경
4. *.go 파일에서 "webserver" 텍스트를 Directory 명으로 변경
6. node_modules 디렉토리 삭제 (디렉토리 이름 변경으로 tailwindcss 오류 발생. 삭제 후 재설치)
7. $ npm i -D tailwindcss
5. site/config.go 파일에서 웹사이트의 기본 정보를 하드코딩
8. store/store.go 파일에서 가게정보들 하드코딩
