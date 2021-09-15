use super::*;

  #[test]
    fn test_system(){ 
        //create new empty system
        let mut system = System::default();

         //add floors to the system
        let mut floors = vec![0,1,2,3];
        system.add_floors(&mut floors);

        //add lifts to the system
        let mut lifts = vec![Lift::new('A', 0, Vec::new(),false,String::from("0"))];
        system.add_lifts(&mut lifts);

        //add calls to the system
        let mut calls = vec![Call::new(3, String::from("Up"))];
        system.add_calls(&mut calls);

        //add requests to the system
        let mut request1 = Request::new('A', 2);
        system.add_requests(&mut request1);
        let mut request2 = Request::new('A', 4);
        system.add_requests(&mut request2);
    
        //test lifts
        assert_eq!('A', system.lifts[0].id);
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
        system.add_lifts(&mut  vec![Lift::new('A', 0, Vec::new(),false,String::from("0"))]);
        system.add_requests(&mut Request::new('A', 2));
        system.move_to_request();
    
        //test lifts
        assert_eq!(2, system.lifts[0].floor);
    }
    #[test]
    fn test_move_down_to_request(){ 
        let mut system = System::default();
        system.add_floors(&mut vec![0,1,2,3]);
        system.add_lifts(&mut  vec![Lift::new('A', 3, vec![1],false,String::from("3"))]);
        system.move_to_request();
    
        //test lifts
        assert_eq!(1, system.lifts[0].floor);
    }
    #[test]
    fn test_move_to_call(){ 
        let mut system = System::default();
        system.add_floors(&mut vec![0,1,2,3]);
        system.add_lifts(&mut  vec![Lift::new('A', 3, vec![1],false,String::from("3"))]);
        system.add_calls(&mut vec![Call::new(1, String::from("Down"))]);
        system.move_to_call();
    
        //test lifts
        assert_eq!(1, system.lifts[0].floor);
    }