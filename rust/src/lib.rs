pub mod print_lifts;

use std::io::prelude::*;
use std::fs::File;

#[derive(Clone)]
struct Lift {
    id: char,
    floor:     u32,
	requests:  Vec<u32>,
	doors_open: bool,
	monitor: String,
}

#[derive(Clone)]
struct Call{
    floor: u32,
    direction: String
}
struct Request {
    lift: char,
    floor: u32
}

#[derive(Clone)]
struct System {
    floors: Vec<u32>,
    lifts:  Vec<Lift>,
    calls:  Vec<Call>
}

impl Lift{
    fn new(id: char, floor: u32, requests: Vec<u32>, doors_open: bool, monitor: String) -> Lift{
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
    fn new(lift: char, floor: u32) -> Request{
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
    fn move_to_call(&mut self){
        let system = self.clone();
        for call in system.calls{
            let mut responder = system.lifts[0].id;
            for lift in &system.lifts{
                if lift.floor == 0 {
                    responder = lift.id
                }
            }
            self.add_requests(&mut Request::new(responder, call.floor));
            self.move_to_request()
        }
    }
  

}

#[cfg(test)]
mod test;