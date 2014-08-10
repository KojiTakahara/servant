var fil = angular.module('numberFilter', []);
fil.filter('padding', function() {
  return function(input) {
    return ("00" + input).slice(-3)
  };
});