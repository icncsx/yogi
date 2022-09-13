namespace go api

// struct SilentClientCount {
// 	// 1: optional i64 count // optional primitive pointer
// 	// 2: optional list<IDCCount> idcCounts, // list of pointer (list itself is literal)
// 	// 3: RoomConfig roomConfig,
// }

struct RoomConfig {
	1: bool multiChatMode
}

// struct IDCCount {
// 	1: string idc, 
// 	2: i64 count,
// }

struct Request {
	1: RoomConfig roomConfig, // always written...
}

struct Response {
	1: i64 statusCode
}

service Echo {
    Response echo(1: Request req)
}