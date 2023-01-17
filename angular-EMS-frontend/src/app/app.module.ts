import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatSidenavModule} from '@angular/material/sidenav'
import {MatButtonModule} from '@angular/material/button'
import {MatIconModule} from '@angular/material/icon'
import {MatDividerModule} from '@angular/material/divider'
import {MatToolbarModule} from '@angular/material/toolbar';
import { ManagerEmployeeListComponent } from './manager-employee-list/manager-employee-list.component'
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ManagerCreateEmployeeComponent } from './manager-create-employee/manager-create-employee.component';
import { FormsModule } from '@angular/forms';
import { ManagerUpdateEmployeeComponent } from './manager-update-employee/manager-update-employee.component';
import { ManagerEmployeeDetailComponent } from './manager-employee-detail/manager-employee-detail.component';
import { ManagerComponent } from './manager/manager.component';
import { ViewPersonalInfoComponent } from './view-personal-info/view-personal-info.component';
import { UpdatePersonalInfoComponent } from './update-personal-info/update-personal-info.component';
import { LeaveComponent } from './leave/leave.component';
import { LeaveManagerComponent } from './leave-manager/leave-manager.component';
import { LeaveEmployeeComponent } from './leave-employee/leave-employee.component';
import { AllemployeesComponent } from './allemployees/allemployees.component';
import { PayrollComponent } from './payroll/payroll.component';
import { DepartmentComponent } from './department/department.component';
import { DepartmentProjectsComponent } from './department-projects/department-projects.component';
import { DepartmentEmployeesComponent } from './department-employees/department-employees.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './logout/logout.component';
import { AuthTokenService } from './auth-token.service';
import { TokenInterceptorService } from './token-interceptor.service';

@NgModule({
  declarations: [
    AppComponent,
    ManagerEmployeeListComponent,
    ManagerCreateEmployeeComponent,
    ManagerUpdateEmployeeComponent,
    ManagerEmployeeDetailComponent,
    ManagerComponent,
    ViewPersonalInfoComponent,
    UpdatePersonalInfoComponent,
    LeaveComponent,
    LeaveManagerComponent,
    LeaveEmployeeComponent,
    AllemployeesComponent,
    PayrollComponent,
    DepartmentComponent,
    DepartmentProjectsComponent,
    DepartmentEmployeesComponent,
    LoginComponent,
    LogoutComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSidenavModule,
    MatButtonModule,
    MatIconModule,
    MatDividerModule,
    MatToolbarModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [{provide:HTTP_INTERCEPTORS , useClass:TokenInterceptorService,multi:true}],
  bootstrap: [AppComponent]
})
export class AppModule { }
