import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllemployeesComponent } from './allemployees/allemployees.component';
import { DepartmentEmployeesComponent } from './department-employees/department-employees.component';
import {  DepartmentProjectsComponent } from './department-projects/department-projects.component';
import { DepartmentComponent } from './department/department.component';
import { LeaveEmployeeComponent } from './leave-employee/leave-employee.component';
import { LeaveManagerComponent } from './leave-manager/leave-manager.component';
import { LeaveComponent } from './leave/leave.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './logout/logout.component';
import { ManagerCreateEmployeeComponent } from './manager-create-employee/manager-create-employee.component';
import { ManagerEmployeeDetailComponent } from './manager-employee-detail/manager-employee-detail.component';
import { ManagerEmployeeListComponent } from './manager-employee-list/manager-employee-list.component';
import { ManagerUpdateEmployeeComponent } from './manager-update-employee/manager-update-employee.component';
import { ManagerComponent } from './manager/manager.component';
import { PayrollComponent } from './payroll/payroll.component';
import { UpdatePersonalInfoComponent } from './update-personal-info/update-personal-info.component';
import { ViewPersonalInfoComponent } from './view-personal-info/view-personal-info.component';


const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'login',
  },
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'logout',
    component: LogoutComponent,
  },
  {
    path: 'manager',
    component: ManagerComponent,
  },
  {
    path: 'employee-list/:eid',
    component: ManagerEmployeeListComponent,
  },
  {
    path: 'create-employee/:eid',
    component: ManagerCreateEmployeeComponent,
  },
  {
    path: 'update-employee/:childEid',
    component: ManagerUpdateEmployeeComponent,
  },
  {
    path: 'employee-detail/:childEid',
    component: ManagerEmployeeDetailComponent,
  },
  {
    path: 'personal/:eid',
    component: ViewPersonalInfoComponent,
  },
  {
    path: 'update-personal/:eid',
    component: UpdatePersonalInfoComponent,
  },
  {
    path: 'leave',
    component: LeaveComponent,
  },
  {
    path: 'leave-manager/:eid',
    component: LeaveManagerComponent,
  },
  {
    path: 'leave-employee/:eid',
    component: LeaveEmployeeComponent,
  },
  {
    path: 'allemployees',
    component: AllemployeesComponent,
  },
  {
    path: 'payroll',
    component: PayrollComponent,
  },
  {
    path: 'department',
    component: DepartmentComponent,
  },
  {
    path: 'department-projects/:did',
    component: DepartmentProjectsComponent,
  } ,
  {
    path: 'department-employees/:did',
    component: DepartmentEmployeesComponent,
  }    
  

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
