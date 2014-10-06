//
// Bootstrap the jQuery select UX widget using 'Selecter.js' by Ben Plum
// http://www.benplum.com/formstone/selecter/
//
// Dependancies:
//   Angular (duh)
//   jQuery (required by selecter.js)
//   Selecter.js
//
//
// On the select element, define these attributes:
//
//    // Attach the directive
//    selecter-for-option-with-ng-repeat
//
//    // Pass in the ID of the select element
//    selecter-target="#linkStartImage"
//
//    // Define a callback function (in this instance it is used
//    // to set the first option as the default)
//    selecter-callback="setSelectedItem"
//
//    // Define any configuration options as an object which
//    // will be passed to the plugin
//    selecter-config="{ defaultLabel: 'foo' }"
//

var selecterForOptionWithNgRepeat = angular.module('selecterForOptionWithNgRepeat', [])
  .directive("selecterForOptionWithNgRepeat", function($timeout, $parse) {

    return function( scope, element, attrs ) {

      // $last & $timeout are a hack to have this run only once, after
      // ng-repeat has finished building out the dom elements.
      //
      // ng-options does not provide $last which is why we are
      // using ng-repeat instead
      if (scope.$last) {
        $timeout(function() {
          var selecterConfig = {};

          jQuery.extend(selecterConfig, $parse(attrs.selecterConfig)());

          jQuery.extend(selecterConfig, {
            callback: function(value, index) {
              // wrap with $apply so angular knows to pay attention
              scope.$apply(function() {
                var propagateF = $parse(attrs.selecterCallback)(scope);
                propagateF(value, index);
              })
            }
          });

          // target the element and initilize selecter() while
          // passing in the config options
          angular.element(attrs.selecterTarget).selecter(selecterConfig);
        });
      }

    }

  });