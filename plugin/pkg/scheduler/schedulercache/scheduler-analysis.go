package schedulercache


NewFromConfigurator()  // 根据Configurator生成scheduler
//NewFromConfigurator returns a new scheduler that is created entirely by the Configurator.  Assumes Create() is implemented.
// Supports intermediate Config mutation for now if you provide modifier functions which will run after Config is created.
scheduler的主要作用
// Scheduler watches for new unscheduled pods. It attempts to find
// nodes that they fit on and writes bindings back to the api server.
addPod()-> calculateResource(pod)计算资源的函数,很有学习意义
+--- .
+--- |   algorithm
+--- |   |   BUILD
+--- |   |   doc.go
+--- |   |   predicates
+--- |   |   |   BUILD
+--- |   |   |   error.go
+--- |   |   |   metadata.go
+--- |   |   |   predicates.go
+--- |   |   |   predicates_test.go
+--- |   |   |   utils.go
+--- |   |   |   utils_test.go
+--- |   |   priorities
+--- |   |   |   balanced_resource_allocation.go
+--- |   |   |   balanced_resource_allocation_test.go
+--- |   |   |   BUILD
+--- |   |   |   image_locality.go
+--- |   |   |   image_locality_test.go
+--- |   |   |   interpod_affinity.go
+--- |   |   |   interpod_affinity_test.go
+--- |   |   |   least_requested.go
+--- |   |   |   least_requested_test.go
+--- |   |   |   metadata.go
+--- |   |   |   metadata_test.go
+--- |   |   |   most_requested.go
+--- |   |   |   most_requested_test.go
+--- |   |   |   node_affinity.go
+--- |   |   |   node_affinity_test.go
+--- |   |   |   node_label.go
+--- |   |   |   node_label_test.go
+--- |   |   |   node_prefer_avoid_pods.go
+--- |   |   |   node_prefer_avoid_pods_test.go
+--- |   |   |   selector_spreading.go
+--- |   |   |   selector_spreading_test.go
+--- |   |   |   taint_toleration.go
+--- |   |   |   taint_toleration_test.go
+--- |   |   |   test_util.go
+--- |   |   |   util
+--- |   |   |   |   BUILD
+--- |   |   |   |   non_zero.go
+--- |   |   |   |   topologies.go
+--- |   |   |   |   topologies_test.go
+--- |   |   |   |   util.go
+--- |   |   scheduler_interface.go
+--- |   |   scheduler_interface_test.go
+--- |   |   types.go
+--- |   |   types_test.go
+--- |   |   well_known_labels.go
+--- |   algorithmprovider
+--- |   |   BUILD
+--- |   |   defaults
+--- |   |   |   BUILD
+--- |   |   |   compatibility_test.go
+--- |   |   |   defaults.go
+--- |   |   |   defaults_test.go
+--- |   |   plugins.go
+--- |   |   plugins_test.go
+--- |   api
+--- |   |   BUILD
+--- |   |   latest
+--- |   |   |   BUILD
+--- |   |   |   latest.go
+--- |   |   register.go
+--- |   |   types.go
+--- |   |   v1
+--- |   |   |   BUILD
+--- |   |   |   register.go
+--- |   |   |   types.go
+--- |   |   validation
+--- |   |   |   BUILD
+--- |   |   |   validation.go
+--- |   |   |   validation_test.go
+--- |   BUILD
+--- |   core
+--- |   |   BUILD
+--- |   |   equivalence_cache.go
+--- |   |   equivalence_cache_test.go
+--- |   |   extender.go
+--- |   |   extender_test.go
+--- |   |   generic_scheduler.go
+--- |   |   generic_scheduler_test.go
+--- |   factory
+--- |   |   BUILD
+--- |   |   factory.go
+--- |   |   factory_test.go
+--- |   |   plugins.go
+--- |   |   plugins_test.go
+--- |   metrics
+--- |   |   BUILD
+--- |   |   metrics.go
+--- |   OWNERS
+--- |   scheduler.go
+--- |   schedulercache
+--- |   |   BUILD
+--- |   |   cache.go
+--- |   |   cache_test.go
+--- |   |   interface.go
+--- |   |   node_info.go
+--- |   |   util.go
+--- |   scheduler_test.go
+--- |   testing
+--- |   |   BUILD
+--- |   |   fake_cache.go
+--- |   |   fake_lister.go
+--- |   |   pods_to_cache.go
+--- |   testutil.go
+--- |   util
+--- |   |   backoff_utils.go
+--- |   |   backoff_utils_test.go
+--- |   |   BUILD
+--- |   |   utils.go


//   +-------------------------------------------+  +----+
//   |                            Add            |  |    |
//   |                                           |  |    | Update
//   +      Assume                Add            v  v    |
//Initial +--------> Assumed +------------+---> Added <--+
//   ^                +   +               |       +
//   |                |   |               |       |
//   |                |   |           Add |       | Remove
//   |                |   |               |       |
//   |                |   |               +       |
//   +----------------+   +-----------> Expired   +----> Deleted
//         Forget             Expire


question:
1.informer的作用?所有资源都有informer pkg/client 目前看所有的资源的controller都会有对应的informer
2.clientset?  pkg/client
3.client-go repo expecially for cache
4. SharedIndexInformer mean what?
5. 如果pod指明了nodename，在调度的时候是否考虑node resource？
6. api latest v1 目录与文件组织？