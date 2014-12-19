var gulp = require('gulp');
var run = require('gulp-run');
var plumber = require('gulp-plumber');
var notify = require('gulp-notify');

gulp.task('test', function() {
    run('go test', {cwd: 'wsconnector'}).exec()
    .on('error', notify.onError({
        title: "Crap",
        message: "Your tests failed, Miraclew!"
    }))
    .pipe(notify({
        title: "Success",
        message: "All tests have returned green!"
    }));
});

gulp.task('watch', function() {
    gulp.watch('wsconnector/**/*.go', ['test']);
});

gulp.task('default', ['test', 'watch']);