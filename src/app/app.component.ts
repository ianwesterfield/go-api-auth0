import 'rxjs/add/operator/take';
import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { AuthService } from './auth/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(public auth: AuthService, private httpClient: HttpClient) {
    auth.handleAuthentication();
  }

  getStatus() {
    this.httpClient.get<any>('http://localhost:3000/api/v1/version').take(1).subscribe(result => alert(JSON.stringify(result)));
  }
}
