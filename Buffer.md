# Buffer

- 두 모듈(소프트웨어 또는 하드웨이) 간 데이터 송수신시 임시 데이터를 보관하는 저장소
- Queue로 되어 있다
- 구현 방법: Slice, `Ring`

## Ring buffer 
- `Ring` : 환형 자료구조
- 맨 뒤와 맨 앞이 연결된 구조
- 고정된 메모리 
- Slice(Queue) 형태는 GC 카운트와 slice의 재할당이 많이 일어나 비효율 발생 