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
}

impl Call{
    fn new(floor: u32, direction: String) -> Call{
        Call{
            floor,
            direction,
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
    
        //test lifts
        assert_eq!("A", system.lifts[0].id);
        assert_eq!(0, system.lifts[0].floor);
        assert_eq!(false, system.lifts[0].doors_open);

        //test calls
        assert_eq!("Up",system.calls[0].direction);
        assert_eq!(3,system.calls[0].floor);

        //test floor
        assert_eq!(vec![0,1,2,3],system.floors);
    }
}