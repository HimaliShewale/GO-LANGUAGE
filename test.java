package myPackage;

import java.sql.Driver;
import java.util.Set;
import java.util.concurrent.TimeUnit;

import org.openqa.selenium.By;
import org.openqa.selenium.NoSuchElementException;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.support.ui.FluentWait;
import org.openqa.selenium.support.ui.Wait;
import org.openqa.selenium.support.ui.WebDriverWait;

public class Test {
	public static void main(String args[]) throws InterruptedException {

		ChromeOptions opt = new ChromeOptions();
		//Enter the path of your Electron app
		opt.setBinary("C:/Program Files/ERA Explorer/ERA Explorer.exe");
		DesiredCapabilities capabilities = new DesiredCapabilities();
		capabilities.setCapability("chromeOptions", opt);
		capabilities.setBrowserName("chrome");
		System.setProperty("webdriver.chrome.driver",
				"C:/Users/himalis/Downloads/chromedriver-v3.1.6-win32-ia32/chromedriver.exe");
		WebDriver driver = new ChromeDriver(capabilities);
		driver.manage().timeouts().implicitlyWait(30, TimeUnit.SECONDS);
		

		//Printing the window title
		System.out.println(driver.getTitle());
		
	}
}