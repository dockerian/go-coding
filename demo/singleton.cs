// ---------------------------------------------------------------------------
// <copyright file="singleton.cs" company="Dockerian">
//   Copyright (c) 2016 Jason Zhu.  All rights reserved.
// </copyright>
// ---------------------------------------------------------------------------

/// <summary>
/// Singleton class demonstrates two ways of creating singleton
/// </summary>
public class Singleton {
  const bool _usingLock = false; // using lock or nested static instance

  // private constructor ensures this cannot be initialized explicitly
  private Singleton() {}

  private static object _syncObj = new object();
  private static Singleton _instance = null;

  // default to `internal` if no class access modifier is specified
  class Nested {
    // default to `private` if no constructor/member access modifier specified
    static Nested() {} // static constructor is called only once by .NET runtime
    // static field is initialized when class constructor is called
    internal static readonly Singleton instance = new Singleton();
  }

  /// <summary>
  /// Singleton.Instance returns singleton instance of Singleton class
  /// </summary>
  public static Singleton Instance {
    get {
      if (_usingLock) {
        if (_instance == null) {
          lock (_syncObj) {
            if (_instance == null)
              _instance = new Singleton();
        }
        return _instance;
      }
      // class `Nested` is constructed when this is called at the 1st time
      return Nested.instance;
    }
  }
}
