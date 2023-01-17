import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EmployeeLeave } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-leave-manager',
  templateUrl: './leave-manager.component.html',
  styleUrls: ['./leave-manager.component.css']
})
export class LeaveManagerComponent implements OnInit {

  employeeLeaves: EmployeeLeave[];
  eid: string; // storing the id of current manager who will aprrove or reject leaves
  constructor(private route: ActivatedRoute,
    private employeeService: EmployeeService,private router:Router) { }

  ngOnInit(): void {
    this.getEmployeeLeaveData()
  }

  private getEmployeeLeaveData(){
    this.eid = this.route.snapshot.params['eid']
    //console.log(this.eid)
    this.employeeService.getEmployeesLeaveData(this.eid).subscribe(data => {
      this.employeeLeaves = data;
    
    });
  }


  // childEid is the id of the employee under manager whose id is eid whose leave request needs to be rejected or aprroved.

  approve(childEid: string){
    this.employeeService.approveLeave(this.eid,childEid).subscribe(data => {
     
    });
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['leave-manager',this.eid])
  }

  reject(childEid: string){
    this.employeeService.rejectLeave(this.eid,childEid).subscribe(data => {
     
    });
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['leave-manager',this.eid])
  }

}
