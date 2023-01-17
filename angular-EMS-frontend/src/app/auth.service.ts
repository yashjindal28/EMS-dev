import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { User } from './employee';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseURL = String(environment.authHost) +"/auth";

  constructor(private httpClient: HttpClient ) { }

  login(user: User){
    return this.httpClient.post(`${this.baseURL}/${'login'}`,user)
     }

  
}
