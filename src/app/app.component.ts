import { AuthService } from '@app/auth/auth.service';
import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { take, map } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  constructor(public authService: AuthService, private httpClient: HttpClient) {
    authService.handleAuthentication();
  }

  sendSupportMessage() {
    this.httpClient.post<any>('http://localhost:3000/api/v1/support', {}).pipe(
      take(1), 
      map(result => { 
        result.body = JSON.parse(result.body); return result;
      })
    ).subscribe(result => {
      console.log(result);
    });
  }
}
