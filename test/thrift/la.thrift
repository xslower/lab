namespace go la
namespace py la

struct Art{
	1:i32 id,
	2:string title,
	3:list<string> content,
}

service kickArt{
	map<i32, string> kick(1:i32 id, 2:string name),
	void put(1: Art na)
}