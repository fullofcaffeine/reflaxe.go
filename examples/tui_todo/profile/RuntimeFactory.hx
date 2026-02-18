package profile;

class RuntimeFactory {
  public static function create():TodoRuntime {
    #if example_profile_gopher
    return new GopherRuntime();
    #elseif example_profile_metal
    return new MetalRuntime();
    #else
    return new PortableRuntime();
    #end
  }
}
