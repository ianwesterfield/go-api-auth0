import { APP_BASE_HREF } from "@angular/common";
import { AppComponent } from "./app.component";
import { AuthService } from "@app/auth/auth.service";
import { RouterModule } from "@angular/router";
import { TestBed, async } from "@angular/core/testing";

describe("AppComponent", () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [RouterModule.forRoot([])],
      providers: [
        { provide: AuthService, useValue: { handleAuthentication: () => {}, isAuthenticated: () => {} } },
        { provide: APP_BASE_HREF, useValue: "/" }
      ],
      declarations: [AppComponent]
    });
    TestBed.compileComponents();
  });

  it("should create the app", async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));
});
