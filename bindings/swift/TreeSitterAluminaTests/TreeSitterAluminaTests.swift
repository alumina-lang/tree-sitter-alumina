import XCTest
import SwiftTreeSitter
import TreeSitterAlumina

final class TreeSitterAluminaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_alumina())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Alumina grammar")
    }
}
