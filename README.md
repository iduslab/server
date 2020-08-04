# Brain Storming Club

브레인스토밍클럽 서버를위한 디스코드봇!

## 명령어

<> : 필수, [] : 옵션 <br />
**모든 명령어 앞에는 prefix를 붙여야합니다**

- 상자추가: 쪽지를 담을 상자를 추가합니다
- 상자리스트: 현제 생성된 상자들의 목록을 보여줍니다
- 쪽지추가 <상자id> <쪽지에 넣을내용>: 쪽지를 상자에 넣습니다
- 상자열기 <상자id> [개수 (기본값=2)]>: 쪽지를 상자에서 몇개 꺼냅니다
- 상자삭제 <상자id>: (관리자전용) 상자와 그 상자에 해당하는 쪽지들을 모두 삭제합니다
- 쪽지리스트 <상자id>: (관리자전용) 상자에 들어있는 모든 쪽지의 내용을 불러옵니다
- 쪽지삭제 <쪽지id>: (관지자전용) 상자에 들어있는 쪽지를 삭제합니다

## 봇 돌리기

config.json폴더를 만들어서 다음과 같이 입력합니다

```json
{
  "token": "봇 토큰",
  "prefix": "봇 prefix",
  "ownerID": "봇 소유자 id"
}
```

다음 명령어를 실행합니다

```bash
# 실행파일로 만들기
go build main.go
```

이제 실행파일을 실행하면 됩니다!<br />
단, config.json파일은 항상 같은 폴더에 들어있어야 합니다

## 라이선스

GPL3.0
