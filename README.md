# webserver
매번 같은 웹서버 코드를 작성하기 때문에 초기화 소스를 미리 작성해둠<br>
모든 광고용 웹서버는 효율적인 유지보수를 위해 해당 소스를 기초로 작성함

1. $ git clone https://jeonghoikun:{{TOKEN}}@github.com/jeonghoikun/webserver.git
2. Directory 이름 변경 (ex. domain.com)
3. $ git config user.name {{USER_NAME}}
4. $ git config user.email {{USER_EMAIL}}
5. $ git remote remove origin
6. $ git remote add origin https://jeonghoikun:{{TOKEN}}@github.com/jeonghoikun/{{DOMAIN}}.git
7. go.mod 파일에서 "webserver" 텍스트를 Directory 명으로 변경
8. *.go 파일에서 "webserver" 텍스트를 Directory 명으로 변경
9. node_modules 디렉토리 삭제 (디렉토리 이름 변경으로 tailwindcss 오류 발생. 삭제 후 재설치)
10. $ npm i -D tailwindcss
11. site/config.go 파일에서 웹사이트의 기본 정보를 하드코딩
12. store/store.go 파일에서 가게정보들 하드코딩
13. static/img/favicon 에서 새로운 파비콘 이미지로 교체
