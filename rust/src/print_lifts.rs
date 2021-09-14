use super::*;
use std::mem::size_of as size_of;

#[derive(Clone)]
struct Printer {

}

impl Printer{
    fn print_lift(self, lift: Lift, floor: u32)-> String{
    let mut liftStr = String::from("");
    let id = lift.id;
    if in_requested_floor(lift, floor) {
		liftStr = format!(" *{}",id);
	} else {
		liftStr = format!("  {}",id);
	}
    return liftStr
    }
}


fn print_lifts(liftSystem: System, liftPrinter: Printer) -> String {
    let mut result = String::from("");
    let floor_number_length = liftSystem.floors.len();
    let floors = liftSystem.floors.clone();
    for floor in floors {
        let calls = print_calls(liftSystem.clone(), floor);
        let call_padding = whiteSpace(2 - calls.len());
		let floor_padding = whiteSpace(floor_number_length-(floor as usize));
		let lifts = print_lifts_for_floor(liftSystem.clone(), liftPrinter.clone(), floor);
        let line = format!("{}{} {}{} {} {}{}\n", 
		floor_padding, floor, calls.join(""), call_padding, lifts.join(" "), floor_padding, floor);
		result.push_str(&line.to_string())
    }
    return result
}


fn print_lifts_for_floor(liftSystem: System, liftPrinter: Printer, floor: u32) -> Vec<String> {
	let mut lifts: Vec<String> = Vec::new();
	for lift in liftSystem.lifts {
		lifts.push(print_lift_for_floor(liftPrinter.clone(), lift, floor))
	}
	return lifts
}

fn print_lift_for_floor(liftPrinter: Printer, lift: Lift, floor: u32) -> String {
    let mut liftStr = String::from("");
	if lift.floor == floor {
		liftStr = liftPrinter.print_lift(lift, floor)
	} else {
		let liftIDPadding = whiteSpace(lift.id as usize);
		if in_requested_floor(lift, floor) {
			liftStr = format!("  *{}",liftIDPadding);
		} else {
			liftStr = format!("   {}",liftIDPadding);
		}
	}
	return liftStr
}

fn in_requested_floor(lift: Lift, floor: u32) -> bool {
    let mut found: bool = false;
	for request in lift.requests {
		if request == floor {
			found = true;
		}
	}
	return found
}

fn print_calls(liftSystem: System, floor: u32) -> Vec<String> {
	let mut calls: Vec<String> = Vec::new();
    let system_calls = calls_for(liftSystem, floor);
	for call in system_calls {
		calls.push(print_call_direction(call));
	}
	return calls;
}

fn calls_for(liftSystem: System, floor: u32) -> Vec<Call> {
	let mut calls: Vec<Call> = Vec::new();
	for c in  liftSystem.calls {
		if c.floor == floor {
			calls.push(c);
		}
	}
	return calls;
}

fn print_call_direction(call: Call)-> String {
	if call.direction == "Down" {
		return String::from("v");
	} else if call.direction == "Up" {
		return String::from("^");
	}
	return String::from(" ")
}

fn whiteSpace(length: usize) -> String {
	let repeated = " ".repeat(length);
    return repeated
}

#[cfg(test)]
mod test{
    use super::*;

    #[test]
    fn print_system(){ 
        let mut system = System::default();
        let printer = Printer{};
        system.add_floors(&mut vec![0,1,2,3]);
        system.add_lifts(&mut  vec![Lift::new('A', 0, Vec::new(),false,String::from("0"))]);
        system.add_requests(&mut Request::new('A', 2));
        assert_eq!("lala", print_lifts(system, printer));
    }
}