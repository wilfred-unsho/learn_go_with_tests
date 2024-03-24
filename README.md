<strong>"# learn_go_with_tests"</strong>
<br>
Go to https://quii.gitbook.io/learn-go-with-tests if you want to be good in Go like me
<br><br>
Basic commands: (1) go run hello.go (2) go mod init hello (3) go test #run test
<br><br>
<em><strong>It is important that your tests are clear specifications of what the code needs to do.</strong><em>
<br><br>
<p>t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier.</p>
