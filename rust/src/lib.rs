struct Lift {
    id: String,
    floor:     u32,
	requests:  Vec<u32>,
	doors_open: bool,
	monitor: String,
}

struct Call{
    floor: u32,
    direction: String
}
struct Request {
    lift: String,
    floor: u32
}

struct System {
    floors: Vec<u32>,
    lifts:  Vec<Lift>,
    calls:  Vec<Call>
}

impl Lift{
    fn new(id: String, floor: u32, requests: Vec<u32>, doors_open: bool, monitor: String) -> Lift{
        Lift{
            id,
            floor,
            requests,
            doors_open,
            monitor,
        }
    }
    fn tick(&mut self, request: u32){
        if request == self.floor{
            self.doors_open = true;
            self.monitor = String::from("*");
        } else if request > self.floor{
            self.floor += 1;
            self.monitor = self.floor.to_string();
        } else if request < self.floor{
            self.floor -=1;
            self.monitor = self.floor.to_string();
        }
    }
}

impl Call{
    fn new(floor: u32, direction: String) -> Call{
        Call{
            floor,
            direction,
        }
    }
}

impl Request{
    fn new(lift: String, floor: u32) -> Request{
        Request{
            lift,
            floor,
        }
    }
}


impl Default for System {
    fn default()-> System {
        System{
            floors: Vec::new(),
            lifts: Vec::new(),
            calls: Vec::new(),
        }
    }
}

impl System{
    fn add_floors(&mut self, floors: &mut Vec<u32>){
        self.floors.append(floors)
    }
    fn add_lifts(&mut self, lifts: &mut Vec<Lift>){
        self.lifts.append(lifts)
    }
    fn add_calls(&mut self, calls: &mut Vec<Call>){
        self.calls.append(calls)
    }
    fn add_requests(&mut self, request: &mut Request){
        for lift in &mut self.lifts{
            if lift.id == request.lift {
                lift.requests.push(request.floor)
            }
        }
    }
    fn move_to_request(&mut self){
         for lift in &mut self.lifts{
            let mut requests = lift.requests.to_vec();
            requests.sort();
            for request in requests{
                while lift.floor != request {
                    lift.tick(request)
                }
            }
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_system(){ 
        //create new empty system
        let mut system = System::default();

         //add floors to the system
        let mut floors = vec![0,1,2,3];
        system.add_floors(&mut floors);

        //add lifts to the system
        let mut lifts = vec![Lift::new(String::from("A"), 0, Vec::new(),false,String::from("0"))];
        system.add_lifts(&mut lifts);

        //add calls to the system
        let mut calls = vec![Call::new(3, String::from("Up"))];
        system.add_calls(&mut calls);

        //add requests to the system
        let mut request1 = Request::new(String::from("A"), 2);
        system.add_requests(&mut request1);
        let mut request2 = Request::new(String::from("A"), 4);
        system.add_requests(&mut request2);
    
        //test lifts
        assert_eq!("A", system.lifts[0].id);
        assert_eq!(0, system.lifts[0].floor);
        assert_eq!(false, system.lifts[0].doors_open);

        //test calls
        assert_eq!("Up",system.calls[0].direction);
        assert_eq!(3,system.calls[0].floor);

        //test floor
        assert_eq!(vec![0,1,2,3],system.floors);

         //test requests
        assert_eq!(vec![2,4], system.lifts[0].requests);

    }
    #[test]
    fn test_move_up_to_request(){ 
        let mut system = System::default();
        system.add_floors(&mut vec![0,1,2,3]);
        system.add_lifts(&mut  vec![Lift::new(String::from("A"), 0, Vec::new(),false,String::from("0"))]);
        system.add_requests(&mut Request::new(String::from("A"), 2));
        system.move_to_request();
    
        //test lifts
        assert_eq!(2, system.lifts[0].floor);
    }
       #[test]
    fn test_move_down_to_request(){ 
        let mut system = System::default();
        system.add_floors(&mut vec![0,1,2,3]);
        system.add_lifts(&mut  vec![Lift::new(String::from("A"), 3, vec![1],false,String::from("3"))]);
        system.move_to_request();
    
        //test lifts
        assert_eq!(1, system.lifts[0].floor);
    }
}