module.exports = function (grunt) {
    var pkg = grunt.file.readJSON('package.json');
    grunt.initConfig({
        concat: {
            files: {
                // 元ファイルの指定
                src : 'static/js/*.js',
                // 出力ファイルの指定
                dest: 'static/js/concat/app.js'
            }
        },

        uglify: {
            dist: {
                files: {
                    // 出力ファイル: 元ファイル
                    'static/js/min/app-min.js': 'static/js/concat/app.js'
                }
            }
        },

        watch: {
            js: {
                files: 'static/js/*.js',
                tasks: ['concat', 'uglify']
            }
        }
    });

    // プラグインのロード・デフォルトタスクの登録
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.registerTask('default', ['concat', 'uglify']);
};
