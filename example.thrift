/*
 * namespace
 */

namespace go thrift.example

/*
 * struct
 */
struct Person {
    1: string name,
    2: i32 age,
}

/*
 * service
 */
service CustomerService {
	list<Person> ListPerson(),
	void AddPerson(1: Person newPerson),
}
